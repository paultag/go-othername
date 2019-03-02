package san

import (
	"encoding/asn1"
)

var (
	// Subject Alternative Name is the ID for the SAN Extension type, in order
	// to extract the SAN specific entries that the Go native bindings don't
	// handle.
	oidSubjectAltName = asn1.ObjectIdentifier{2, 5, 29, 17}

	// The userPrincipalName attribute is the logon name for the user. The
	// attribute consists of a user principal name (UPN), which is the most
	// common logon name for Windows users. Users typically use their UPN to
	// log on to a domain. This attribute is an indexed string that is
	// single-valued.
	oidUPN = asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 311, 20, 2, 3}
)
