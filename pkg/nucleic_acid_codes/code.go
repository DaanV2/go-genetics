package nucleic_acids_codes

import (
	"errors"
	"fmt"
)

// Nucleic Acid Code
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

// Checks if the code is an valid nucleic acid code
func (code Code) Valid() bool {
	return Valid(code)
}

// Returns the associated description of the code
func (code Code) Description() string {
	return Description(code)
}

// From converts the given byte into a [Code]. NOTE: this does no checking of validity of the data. Use [Parse] or [ParseByte] to have it return an error if its invalid
// Else you can also do [Code.Valid()]
//
// Example:
//
//	c := From('M')
//	if c.Valid() {
//		...
//	}
func From(a byte) Code {
	return Code(a)
}

func Parse(code string) (Code, error) {
	if len(code) < 1 {
		return Any, errors.New("expected atleast 1 character to read")
	}

	return ParseByte(code[0])
}

func ParseByte(a byte) (Code, error) {
	c := From(a)
	if Valid(c) {
		return c, nil
	}

	return c, fmt.Errorf("unknown nucleic acid code: '%d' (%v)", a, a)
}
