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

package othername // import "pault.ag/go/othername"

import (
	"crypto/x509"
	"encoding/asn1"
	"fmt"

	"pault.ag/go/othername/fasc"
)

// Return all FASC entries from this Certificate, parsed into a fasc.FASC
// struct.
func FASCs(cert *x509.Certificate) ([]fasc.FASC, error) {
	ons, err := All(cert)
	if err != nil {
		return nil, err
	}
	return ons.FASCs()
}

// Get all FASC entries in the set of OtherNames this function hangs off of.
func (on OtherNames) FASCs() ([]fasc.FASC, error) {
	ret := []fasc.FASC{}

	err := on.Find(oidFASCN).Map(func(on OtherName) error {
		name, err := on.FASC()
		if err != nil {
			return err
		}
		ret = append(ret, *name)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return ret, nil
}

// Decode and parse the FASC (Federal Agency Smartcard Number). This contains
// some basic information on the PIV token that this Certificate belongs to.
func (on OtherName) FASC() (*fasc.FASC, error) {
	if !on.Id.Equal(oidFASCN) {
		return nil, InvalidOID
	}

	rv := asn1.RawValue{}
	bytes, err := on.Unmarshal(&rv)
	if err != nil {
		return nil, err
	}

	if len(bytes) != 0 {
		return nil, fmt.Errorf("othername: OtherName.UPN: Trailing bytes")
	}

	return fasc.Parse(rv.Bytes)
}

// vim: foldmethod=marker
