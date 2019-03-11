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

package fasc // import "pault.ag/go/othername/fasc"

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
	groups, err := group(numbers)
	if err != nil {
		return nil, err
	}
	return unpack(groups)
}

// Extract FASC int streams into chunks based on field seperrators.
func group(values []int) ([][]int, error) {
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

// unpack the list of lists into a FASC Struct.
func unpack(entries [][]int) (*FASC, error) {
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
	ret.OrganizationIdentifier = AgencyCode(rehydrateNumber(rest[11:15]))
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
	// Agency Code identifies what Government agnecy issued the credential.
	// This is usually a FIPS 95-2 federal agency code. An incomplete table
	// of these values is provided in this library, but this type is just a
	// retyped Int, direct access of the underlying value is encouraged.
	AgencyCode AgencyCode

	// Identifies the system the card is enrolled in.
	SystemCode int

	// Unique for the given system. This plus the system and agency code should
	// be a unique identifier for this token.
	Credential int

	// Major version changes in the data above.
	CredentialSeries int

	// Individual incrementing counter for the number of times a card has been
	// replaced due to loss or damage.
	IndidvidualCredentialSeries int

	// Unique identifier for an individual. This is sometimes a fairly low
	// incrementing identifier for the ACL system, and other times it's an
	// agency-wide globally unique identifier such as an EDIPI.
	PersonIdentifier int

	// Type of orgnaization that the cardholder belongs to.
	//
	// A table of values is provided in this library, but this is a masked
	// integer in the event you need direct access. It's likely easier to
	// set up a switch statement using the defined constants on this, though.
	//
	// Valid values here are:
	//
	//  - Federal Government Agency
	//  - State Government Agency
	//  - Commercial Enterprise
	//  - Foreign Government
	//
	OrganizationCategory OrganizationalCategory

	// For the particular category above, this is an iditifier to determine
	// which agency, state, company or government this cardholder is a part
	// of.
	//
	// Given a Federal Government Agency, this will be a FIPS 95-2 agency code.
	// Given a State Government Agency, this will be a State Code.
	// Given a Commercial Enterprise, this will be a Company Code
	// Given a Foreign Government, this will be a Country Code.
	OrganizationIdentifier AgencyCode

	// What the relation of this person to the above organization is.
	//
	// This is a masked int, and direct access to the underlying value may
	// be helpful, but it may be more helpful to use the defined constants.
	//
	// Valid values here are:
	//
	//  - Employee
	//  - Civil
	//  - Executive Staff
	//  - Uniformed Service
	//  - Contractor
	//  - Organizational Affiliate
	//  - Organizational Beneficiary
	//
	PersonAssociation AssociationCategory
}

func (f FASC) String() string {
	return fmt.Sprintf(
		"agency=%s system=%d credential=%d credentialSeries=%d ics=%d pi=%d oc=%s oi=%s assoc=%s",
		f.AgencyCode.String(),
		f.SystemCode,
		f.Credential,
		f.CredentialSeries,
		f.IndidvidualCredentialSeries,
		f.PersonIdentifier,
		f.OrganizationCategory.String(),
		f.OrganizationIdentifier.String(),
		f.PersonAssociation.String(),
	)
}

// vim: foldmethod=marker
