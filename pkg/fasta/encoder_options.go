package fasta

type EncoderOption interface {
	apply(e *Encoder) error
}

type encoderOptionFN struct {
	fn func(e *Encoder) error
}

func (o encoderOptionFN) apply(e *Encoder) error {
	return o.fn(e)
}
func newEncoderOption(fn func(e *Encoder) error) encoderOptionFN {
	return encoderOptionFN{fn}
}

// WithPrintWidth sets the print width for the data, each line should be no longer that this value.
// By default its 80
func WithPrintWidth(width int) EncoderOption {
	return newEncoderOption(func(e *Encoder) error {
		e.printWidth = width
		return nil
	})
}

// WithEndingNewline turns on or off the options to add new lines add the end of the sequence.
// Defaults to true.
func WithEndingNewline(add bool) EncoderOption {
	return newEncoderOption(func(e *Encoder) error {
		e.endWithNewLine = add
		return nil
	})
}

// WithEndingAstrik turns on or off the options to a '*' add the end of the sequence.
// Defaults to false.
func WithEndingAstrik(add bool) EncoderOption {
	return newEncoderOption(func(e *Encoder) error {
		e.endWithAsterik = add
		return nil
	})
}
