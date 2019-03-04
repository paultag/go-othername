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
	"encoding/asn1"
	"fmt"

	"pault.ag/go/othername/fasc"
)

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

	// fasc, err := fasc.Parse([]byte{
	// 	0xd0, 0x43, 0x94, 0x58, 0x21, 0xc, 0x2c, 0x19, 0xa0, 0x84, 0x6d, 0x83,
	// 	0x68, 0x5a, 0x10, 0x82, 0x10, 0x8c, 0xe7, 0x39, 0x84, 0x10, 0x8c, 0xa3,
	// 	0xf5,
	// })
	return fasc.Parse(rv.Bytes)
}

// vim: foldmethod=marker
