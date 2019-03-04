package fasc_test

import (
	"fmt"
	"io"
	"log"
	"testing"

	"pault.ag/go/othername/fasc"
)

/*
 *
 */

func isok(t *testing.T, err error) {
	if err != nil && err != io.EOF {
		log.Printf("Error! Error is not nil! %s\n", err)
		t.FailNow()
	}
}

func assert(t *testing.T, expr bool) {
	if !expr {
		log.Printf("Assertion failed!")
		t.FailNow()
	}
}

func TestExampleEntry(t *testing.T) {
	/* This example was taken from a federal standard PDF, and the byte
	 * blobs below should match the following values:
	 *
	 * AGENCY CODE = 0032
	 * SYSTEM CODE = 0001
	 * CREDENTIAL# = 092446
	 * CS = 0
	 * ICI = 1
	 * PI = 1112223333
	 * OC= 1
	 * OI=1223
	 * POA=2
	 * LRC = 5 */

	f, err := fasc.Parse([]byte{
		0xd0, 0x43, 0x94, 0x58, 0x21, 0xc, 0x2c, 0x19, 0xa0, 0x84, 0x6d, 0x83,
		0x68, 0x5a, 0x10, 0x82, 0x10, 0x8c, 0xe7, 0x39, 0x84, 0x10, 0x8c, 0xa3,
		0xf5,
	})
	isok(t, err)

	assert(t, f.AgencyCode == 32)
	assert(t, f.SystemCode == 1)
	assert(t, f.Credential == 92446)
	assert(t, f.CredentialSeries == 0)
	assert(t, f.IndidvidualCredentialSeries == 1)
	assert(t, f.PersonIdentifier == 1112223333)
	assert(t, f.OrganizationCategory == fasc.OrganizationalCategoryFederalGoverment)
	assert(t, f.OrganizationIdentifier == 1223)
	assert(t, f.PersonAssociation == fasc.AssociationCategoryCivil)
}
