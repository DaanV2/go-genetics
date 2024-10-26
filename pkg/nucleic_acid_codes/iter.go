package nucleic_acids_codes

import "iter"

// Iter returns an iterator over all the characters in the string, do note, the code may be invalid. use [Code.Valid()] to check.
func Iter(data string) iter.Seq2[int, Code] {
	return IterBytes([]byte(data))
}

// IterBytes returns an iterator over all the characters in the slice of bytes, do note, the code may be invalid. use [Code.Valid()] to check.
func IterBytes(data []byte) iter.Seq2[int, Code] {
	return func(yield func(int, Code) bool) {
		for i, b := range data {
			if !yield(i, From(b)) {
				return
			}
		}
	}
}
