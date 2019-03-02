package upn

import (
	"fmt"

	"encoding/asn1"
)

//
type OtherName struct {
	Id    asn1.ObjectIdentifier
	Value asn1.RawValue
}

//
func ParseOtherNames(bytes []byte) ([]OtherName, error) {
	// Given the SubjectAltName Extension bytes, go through the SubjectAltNames,
	// and pull all OtherName SANs out.

	ret := []OtherName{}

	values := []asn1.RawValue{}
	_, err := asn1.Unmarshal(bytes, &values)
	if err != nil {
		return nil, err
	}

	for _, value := range values {
		switch value.Tag {
		case 0: // OtherName
			on, err := ParseOtherName(value.Bytes)
			if err != nil {
				return nil, err
			}
			ret = append(ret, *on)
		}
	}

	return ret, nil
}

//
func ParseOtherName(bytes []byte) (*OtherName, error) {
	var err error

	// OtherName is a encoded blob that contains two ASN1 encoded objects,
	// an ObjectIdentifier, and a RawValue. For the case of a UPN, the
	// nested RawValue is a Unicode String. This will defer the parsing of
	// UPN specific data until we check the ID is right.

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
		return nil, fmt.Errorf("upn: other name contains trailing bytes")
	}

	return &OtherName{
		Id:    id,
		Value: rv,
	}, nil
}
