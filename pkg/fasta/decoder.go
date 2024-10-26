package fasta

import (
	"bufio"
	"errors"
	"io"
)

type Decoder struct {
}

func NewDecoder() Decoder {
	return Decoder{}
}

func (d *Decoder) Decode(reader io.Reader) ([]*Sequence, error) {
	r := FastaReader{
		streamReader: bufio.NewReader(reader),
		sequences:    make([]*Sequence, 0),
		current:      new(Sequence),
	}

	err := r.DecodeAll()
	return r.Sequences(), err
}

type LineType int

const (
	EmptyLine LineType = iota
	DataLine
	DescriptionLine
	CommentLine
	EOF = 4
)

type FastaReader struct {
	streamReader *bufio.Reader
	sequences    []*Sequence
	current      *Sequence
}

func (reader *FastaReader) DecodeAll() error {
	for {
		str, lineType, err := reader.ReadLine()
		if err != nil {
			return NewSequenceError(reader.current, err)
		}

		switch lineType {
		case EOF, EmptyLine: // Marks the end of the sequence / file
			reader.EndCurrent()
			return nil
		case DescriptionLine:
			reader.EndCurrent() // Checks if we need to end anything
			reader.NewDescription(str)
		case CommentLine:
			reader.AddComment(str)
		case DataLine:
			reader.AddData(str)
		}
	}
}

func (reader *FastaReader) Sequences() []*Sequence {
	return reader.sequences
}

func (reader *FastaReader) ReadLine() (string, LineType, error) {
	str, err := reader.streamReader.ReadString('\n')
	if err != nil {
		if errors.Is(err, io.EOF) {
			return str, EOF, nil
		}
		return str, EmptyLine, err
	}

	switch str[0] {
	case '>':
		return Trim(str), DescriptionLine, nil
	case ':':
		return Trim(str), CommentLine, nil
	}

	return str, DataLine, nil
}

func (reader *FastaReader) NewDescription(line string) {
	reader.current.Description = line
}

func (reader *FastaReader) AddComment(line string) {
	reader.current.Comments = append(reader.current.Comments, line)
}

func (reader *FastaReader) AddData(line string) {
	reader.current.Description += line
}

func (reader *FastaReader) EndCurrent() {
	if reader.current == nil || len(reader.current.Description) > 0 {
		if reader.current != nil {
			reader.sequences = append(reader.sequences, reader.current)
		}

		reader.current = new(Sequence)
	}
}
