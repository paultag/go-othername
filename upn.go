package san

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
