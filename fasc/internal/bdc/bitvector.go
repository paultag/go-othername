package bdc

// A bitslice is an array of bits. This allows us to slice and re-interpret
// the bits, regardless of how they got to us.
//
// For AIS, we get data in 6-bit bytes, but we need to slice them up without
// regard to any particular alignment.
//
// This is a memory hog (it's a slice of bools), and not very fast, but it's
// going to be a good place to start optimizing from, once it's all working.
type BitSlice struct {
	bits []bool
}

// Get the number of bits in this bit vector.
func (bv *BitSlice) Length() int {
	return len(bv.bits)
}

// Append a single bit to the bit slice.
func (bv *BitSlice) Append(bit bool) {
	bv.bits = append(bv.bits, bit)
}

// Slice the bit slice from bit offset `from`, for `length` bits. That slice
// is returned as another BitSlice, which can be interpreted from there.
func (bv *BitSlice) Slice(from, length uint) *BitSlice {
	if int(from+length) > len(bv.bits) {
		return &BitSlice{bits: []bool{}}
	}

	return &BitSlice{
		bits: bv.bits[from : from+length],
	}
}

// Read the bit slice as a uint. This assumes the Nth index is the MSB.
func (bv *BitSlice) BigEndianUint() uint {
	ret := uint(0)
	for i := range bv.bits {
		if bv.bits[i] {
			ret += (0x01 << uint(i))
		}
	}
	return ret
}

// Append a byte to the list, MSB first.
func (bv *BitSlice) AppendByte(b byte, length uint) {
	for i := int(length - 1); i >= 0; i-- {
		bv.Append((b & (0x01 << uint(i))) != 0)
	}
}

// Append a slice of bytes.
func (bv *BitSlice) AppendBytes(data []byte, length uint) {
	for _, el := range data {
		bv.AppendByte(el, length)
	}
}

// Allocate a new BitSlice.
func NewBitSlice() BitSlice {
	return BitSlice{bits: []bool{}}
}

//
func NewBitSliceFromBytes(data []byte) (*BitSlice, error) {
	b := NewBitSlice()
	for _, el := range data {
		b.AppendByte(el, 8)
	}
	return &b, nil
}
