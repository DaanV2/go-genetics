package xstrings

import (
	"io"
	"unicode/utf8"
)

type Writer struct {
	base io.Writer
}

func NewWriter(writer io.Writer) Writer {
	return Writer{writer}
}

func (w *Writer) WriteString(data string) error {
	_, err := w.base.Write([]byte(data))
	return err
}

func (w *Writer) WriteRune(data rune) error {
	// Runes are at most 4 bytes of encoding
	b := [4]byte{}

	n := utf8.EncodeRune(b[:], data)
	_, err := w.base.Write(b[:n])
	return err
}
