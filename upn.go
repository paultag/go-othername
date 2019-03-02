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

package othername

import (
	"fmt"

	"crypto/x509"
	"encoding/asn1"
)

// Extract the Universal Principal Name
func UPNs(cert *x509.Certificate) ([]string, error) {
	ons, err := ParseOtherNames(cert)
	if err != nil {
		return nil, err
	}
	return ons.UPNs()
}

//
//
//
func (on OtherName) UPN() (string, error) {
	if !on.Id.Equal(oidUPN) {
		return "", fmt.Errorf("san: OtherName.UPN: Wrong ObjectIdentifier for a UPN")
	}

	bytes := on.Value.Bytes
	upn := asn1.RawValue{}
	bytes, err := asn1.Unmarshal(bytes, &upn)
	if err != nil {
		return "", err
	}

	if len(bytes) != 0 {
		return "", fmt.Errorf("san: OtherName.UPN: Trailing bytes")
	}

	return string(upn.Bytes), nil
}

func (on OtherNames) UPNs() ([]string, error) {
	ret := []string{}
	upns := on.Find(oidUPN)
	for _, upn := range upns {
		name, err := upn.UPN()
		if err != nil {
			return nil, err
		}
		ret = append(ret, name)
	}
	return ret, nil
}

// vim: foldmethod=marker
