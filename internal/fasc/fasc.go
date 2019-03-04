package fasc

import (
	"fmt"

	"pault.ag/go/othername/internal/bdc"
)

func Parse(data []byte) (*FASC, error) {
	numbers, err := bdc.Parse(data)
	if err != nil {
		return nil, err
	}
	groups, err := Group(numbers)
	if err != nil {
		return nil, err
	}
	return Unpack(groups)
}

func Group(values []int) ([][]int, error) {
	if values[0] != 11 {
		return nil, fmt.Errorf("fasc: start byte isn't 1101")
	}

	ret := [][]int{}
	current := []int{}

	for _, value := range values[1:] {
		if value == 13 {
			ret = append(ret, current)
			current = []int{}
			continue
		}
		current = append(current, value)
	}

	ret = append(ret, current)
	return ret, nil
}

// SS AGENCY CODE + FS SYSTEM CODE + FS CREDENTIAL NUMBER + FS CS + FS ICI + FS
// PI + OC + OI + POA + ES + LRC

func getElement(entries [][]int, index int, length int) []int {
	ret := entries[index]
	if len(ret) != length {
		return nil
	}
	return ret
}

// very very very very fragle. this needs serious help
func Unpack(entries [][]int) (*FASC, error) {
	ret := FASC{
		AgencyCode:                  getElement(entries, 0, 4),
		SystemCode:                  getElement(entries, 1, 4),
		Credential:                  getElement(entries, 2, 6),
		CredentialSeries:            getElement(entries, 3, 1)[0],
		IndidvidualCredentialSeries: getElement(entries, 4, 1)[0],
	}

	rest := entries[5]
	ret.PersonIdentifier = rest[:10]
	ret.OrganizationCategory = OrganizationalCategory(rest[10])
	ret.OrganizationIdentifier = rest[11:15]
	ret.PersonAssociation = AssociationCategory(rest[15])

	return &ret, nil
}

type FASC struct {
	AgencyCode                  []int
	SystemCode                  []int
	Credential                  []int
	CredentialSeries            int
	IndidvidualCredentialSeries int

	PersonIdentifier       []int
	OrganizationCategory   OrganizationalCategory
	OrganizationIdentifier []int
	PersonAssociation      AssociationCategory
}

func (f FASC) String() string {
	return fmt.Sprintf(
		"agency=%d system=%d credential=%d credentialSeries=%d ics=%d pi=%d oc=%s oi=%d assoc=%s",
		f.AgencyCode,
		f.SystemCode,
		f.Credential,
		f.CredentialSeries,
		f.IndidvidualCredentialSeries,
		f.PersonIdentifier,
		f.OrganizationCategory.String(),
		f.OrganizationIdentifier,
		f.PersonAssociation.String(),
	)
}
