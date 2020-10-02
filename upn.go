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
	"fmt"

	"crypto/x509"
	"encoding/asn1"
)

// UPNs will extract all Universal Principal Names from an x509 Certificate.
func UPNs(cert *x509.Certificate) ([]string, error) {
	ons, err := All(cert)
	if err != nil {
		return nil, err
	}
	return ons.UPNs()
}

// UPN will extract and return the Universal Principal Name.
func (on OtherName) UPN() (string, error) {
	if !on.ID.Equal(oidUPN) {
		return "", ErrInvalidOID
	}

	upn := asn1.RawValue{}
	bytes, err := on.Unmarshal(&upn)
	if err != nil {
		return "", err
	}

	if len(bytes) != 0 {
		return "", fmt.Errorf("othername: OtherName.UPN: Trailing bytes")
	}

	return string(upn.Bytes), nil
}

// UPNs will extract and return all Universal Principal Names from a list of
// OtherNames.
func (on OtherNames) UPNs() ([]string, error) {
	ret := []string{}

	err := on.Find(oidUPN).Map(func(on OtherName) error {
		name, err := on.UPN()
		if err != nil {
			return err
		}
		ret = append(ret, name)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return ret, nil
}

// vim: foldmethod=marker
