package xerrors

import (
	"errors"
	"fmt"
)

type IndexedError struct {
	index int
	err   error
}

func NewIndexedError(index int, err error) *IndexedError {
	return &IndexedError{
		index,
		err,
	}
}

func (s *IndexedError) Error() string {
	return fmt.Errorf("error caused by item in collection[%v]: %w", s.index, s.err).Error()
}

func (s *IndexedError) Is(err error) bool {
	return errors.Is(s.err, err)
}

func (s *IndexedError) Unwrap() error {
	return s.err
}
