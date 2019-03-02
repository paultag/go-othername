package san

import (
	"fmt"

	"encoding/asn1"
)

// Encapsulation of an x509 Subject Alternative Name (SAN) Other Name.
//
// This contains an ObjectIdentifier Id for the OtherName type, and the
// opaque type dependent RawValue.
type OtherName struct {
	Id    asn1.ObjectIdentifier
	Value asn1.RawValue
}

type OtherNames []OtherName

func (o OtherNames) Find(id asn1.ObjectIdentifier) OtherNames {
	ret := OtherNames{}
	for _, on := range o {
		if on.Id.Equal(id) {
			ret = append(ret, on)
		}
	}
	return ret
}

//
//
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
