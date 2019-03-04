package bdc

import (
	"pault.ag/go/othername/internal/bdc/bitvector"
)

// Interpret the input bytes as a BDC 4-bit number stream.
func Parse(bytes []byte) ([]int, error) {
	bv, err := bitvector.Decode(bytes)
	if err != nil {
		return nil, err
	}

	ret := []int{}
	length := bv.Length()

	for i := 0; i < length; i = i + 5 {
		ret = append(ret, int(bv.Slice(uint(i), 4).BigEndianUint()))
	}

	return ret, nil
}
