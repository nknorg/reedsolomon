package reedsolomon

// BytesArray is a gomobile compatible wrapper of [][]byte.
type BytesArray struct{ elems [][]byte }

// NewBytesArray creates an empty BytesArray with length n.
func NewBytesArray(n int) *BytesArray {
	return &BytesArray{
		elems: make([][]byte, n),
	}
}

// Len returns the bytes array length.
func (ba *BytesArray) Len() int {
	return len(ba.elems)
}

// Get returns an element of the BytesArray.
func (ba *BytesArray) Get(i int) []byte {
	return ba.elems[i]
}

// Set sets an element of the BytesArray.
func (ba *BytesArray) Set(i int, b []byte) {
	ba.elems[i] = b
}

// Append appends an element to the BytesArray and increments its length.
func (ba *BytesArray) Append(b []byte) {
	ba.elems = append(ba.elems, b)
}
