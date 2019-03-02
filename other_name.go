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
	"crypto/x509"
	"encoding/asn1"
	"fmt"
)

// Encapsulation of an x509 Subject Alternative Name (SAN) Other Name.
//
// This contains an ObjectIdentifier Id for the OtherName type, and the
// opaque type dependent RawValue. This type is pretty hard to work with
// and will likely involve some calls to asn1.Unmarshal to extract the
// meaningful data, so be sure you feel comfortable with that before
// digging into Value!
type OtherName struct {
	// ObjectIdentifier defining what type of information is contained
	// inside the Value object. There's basically no reason to read the
	// Value without first checking the ObjectIdentifier.
	Id asn1.ObjectIdentifier

	// Type-specific information regarding this particular OtherName.
	Value asn1.RawValue
}

// Unmarshal the payload bytes inside the Value into an interface.
func (o OtherName) Unmarshal(target interface{}) ([]byte, error) {
	return asn1.Unmarshal(o.Value.Bytes, target)
}

//
type OtherNames []OtherName

// Find all OtherNames that have the ObjectIdentifier provided.
func (o OtherNames) Find(id asn1.ObjectIdentifier) OtherNames {
	ret := OtherNames{}
	for _, on := range o {
		if on.Id.Equal(id) {
			ret = append(ret, on)
		}
	}
	return ret
}

// Given a Certificate, go through all the Extensions, find the SubjectAltName
// Extension, and extract all OtherNames from the SAN.
func ParseOtherNames(cert *x509.Certificate) (OtherNames, error) {
	for _, extension := range cert.Extensions {
		if extension.Id.Equal(oidSubjectAltName) {
			ons, err := otherNamesFromSANBytes(extension.Value)
			if err != nil {
				return nil, err
			}
			return ons, nil
		}
	}
	return OtherNames{}, nil
}

// Given the SubjectAltName Extension bytes, go through the SubjectAltNames,
// and pull all OtherName SANs out.
func otherNamesFromSANBytes(bytes []byte) (OtherNames, error) {

	ret := OtherNames{}

	values := []asn1.RawValue{}
	_, err := asn1.Unmarshal(bytes, &values)
	if err != nil {
		return nil, err
	}

	for _, value := range values {
		switch value.Tag {
		case 0: // OtherName
			on, err := otherNameFromBytes(value.Bytes)
			if err != nil {
				return nil, err
			}
			ret = append(ret, *on)
		}
	}

	return ret, nil
}

// OtherName is a encoded blob that contains two ASN1 encoded objects,
// an ObjectIdentifier, and a RawValue. For the case of a UPN, the
// nested RawValue is a Unicode String. This will defer the parsing of
// UPN specific data until we check the ID is right.
func otherNameFromBytes(bytes []byte) (*OtherName, error) {
	var err error

	id := asn1.ObjectIdentifier{}
	rv := asn1.RawValue{}

	// First object is the ObjectIdentifier.
	bytes, err = asn1.Unmarshal(bytes, &id)
	if err != nil {
		return nil, err
	}

	// Next is the opaque ID specific blob.
	bytes, err = asn1.Unmarshal(bytes, &rv)
	if err != nil {
		return nil, err
	}

	if len(bytes) != 0 {
		return nil, fmt.Errorf("san: other name contains trailing bytes")
	}

	return &OtherName{
		Id:    id,
		Value: rv,
	}, nil
}

// vim: foldmethod=marker
