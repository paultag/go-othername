// {{{ Copyright (c) Paul R. Tagliamonte <paultag@gmail.com>, 2019
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE. }}}

package fasc

import (
	"fmt"

	"pault.ag/go/othername/fasc/internal/bdc"
)

// Parse packed 8-bit-byte FASC blobs into a FASC struct containing the parsed
// information.
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

// Extract FASC int streams into chunks based on field seperrators.
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

func getElement(entries [][]int, index int, length int) []int {
	ret := entries[index]
	if len(ret) != length {
		return nil
	}
	return ret
}

// Unpack the list of lists into a FASC Struct.
func Unpack(entries [][]int) (*FASC, error) {
	ret := FASC{
		AgencyCode:                  AgencyCode(rehydrateNumber(getElement(entries, 0, 4))),
		SystemCode:                  rehydrateNumber(getElement(entries, 1, 4)),
		Credential:                  rehydrateNumber(getElement(entries, 2, 6)),
		CredentialSeries:            getElement(entries, 3, 1)[0],
		IndidvidualCredentialSeries: getElement(entries, 4, 1)[0],
	}

	rest := entries[5]
	ret.PersonIdentifier = rehydrateNumber(rest[:10])
	ret.OrganizationCategory = OrganizationalCategory(rest[10])
	ret.OrganizationIdentifier = rehydrateNumber(rest[11:15])
	ret.PersonAssociation = AssociationCategory(rest[15])

	return &ret, nil
}

func rehydrateNumber(in []int) int {
	ret := 0
	for _, el := range in {
		ret *= 10
		ret = ret + el
	}
	return ret
}

type FASC struct {
	AgencyCode                  AgencyCode
	SystemCode                  int
	Credential                  int
	CredentialSeries            int
	IndidvidualCredentialSeries int

	PersonIdentifier       int
	OrganizationCategory   OrganizationalCategory
	OrganizationIdentifier int
	PersonAssociation      AssociationCategory
}

func (f FASC) String() string {
	return fmt.Sprintf(
		"agency=%s system=%d credential=%d credentialSeries=%d ics=%d pi=%d oc=%s oi=%d assoc=%s",
		f.AgencyCode.String(),
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

// vim: foldmethod=marker
