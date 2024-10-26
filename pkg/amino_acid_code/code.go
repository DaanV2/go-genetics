package amino_acid_codes

// Amino Acid Code
type Code byte

// Data returns the byte value of the [Code]
func (code Code) Data() uint8 {
	return uint8(code)
}

// String returns the string representation of this code
func (code Code) String() string {
	return string(code)
}

// GoString returns the string representation of this code
func (code Code) GoString() string {
	return code.String()
}

// Checks if the code is an valid amino acid code
func (code Code) Valid() bool {
	return Valid(code)
}

// Returns the associated description of the code
func (code Code) Description() string {
	return Description(code)
}
