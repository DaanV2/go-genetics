package fasta

import (
	"io"

	xstrings "github.com/DaanV2/go-genetics/pkg/strings"
)

type Encoder struct {
	printWidth     int
	endWithNewLine bool // Seperate each sequence with an empty line
	endWithAsterik bool // End the sequence with an asterik
}

func NewEncoder() *Encoder {
	return &Encoder{
		printWidth:     80,
		endWithNewLine: true,
	}
}

func (e *Encoder) Encode(writer io.Writer, sequences ...*Sequence) error {
	for _, seq := range sequences {
		if err := e.encodeItem(writer, seq); err != nil {
			return NewSequenceError(seq, err)
		}
	}

	return nil
}

func (e *Encoder) encodeItem(writer io.Writer, seq *Sequence) error {
	bwriter := xstrings.NewWriter(writer)
	var err error

	// Write the comments as ;
	for _, c := range seq.Comments {
		err = bwriter.WriteString("; " + c)
		if err != nil {
			return err
		}
	}

	// Write description
	err = bwriter.WriteString(">" + seq.Description)
	if err != nil {
		return err
	}

	// Write the data
	data := seq.Data
	for len(data) > 0 {
		// Chunk the data into sets of n length
		maxWrite := min(len(data), e.printWidth)
		s := data[:maxWrite]
		data = data[maxWrite:]

		// Write the data
		err = bwriter.WriteString(s)
		if err != nil {
			return err
		}
		// Have we reached the end? put a * at the end
		if len(data) == 0 && e.endWithAsterik {
			if err := bwriter.WriteString("*"); err != nil {
				return err
			}
		}
		err = bwriter.WriteString("\n")
		if err != nil {
			return err
		}

	}

	// Write empty line
	if e.endWithNewLine {
		err = bwriter.WriteString("\n")
		if err != nil {
			return err
		}
	}

	return nil
}
