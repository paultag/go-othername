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

// vim: foldmethod=marker
