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

	// FASC or Federal Agency Smartcard Number, is information relating to the
	// PIV cardholder.
	oidFASCN = asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 6, 6}
)

// vim: foldmethod=marker
