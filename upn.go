package upn

import (
	"crypto/x509"
	"encoding/asn1"
	"fmt"
)

//
func From(cert *x509.Certificate) ([]string, error) {
	ret := []string{}

	for _, extension := range cert.Extensions {
		if extension.Id.Equal(oidSubjectAltName) {
			ons, err := ParseOtherNames(extension.Value)
			if err != nil {
				return nil, err
			}

			for _, on := range ons {
				if on.Id.Equal(oidUPN) {
					name, err := ParseUPN(on)
					if err != nil {
						return nil, err
					}
					ret = append(ret, name)
				}
			}
		}
	}
	return ret, nil
}

func ParseUPN(on OtherName) (string, error) {
	if !on.Id.Equal(oidUPN) {
		return "", fmt.Errorf("other name upn broken")
	}

	bytes := on.Value.Bytes
	upn := asn1.RawValue{}
	bytes, err := asn1.Unmarshal(bytes, &upn)
	if err != nil {
		return "", err
	}

	if len(bytes) != 0 {
		return "", fmt.Errorf("other name short bytes")
	}

	return string(upn.Bytes), nil
}
