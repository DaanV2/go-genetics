package fasta

import (
	"errors"
	"fmt"
)

var _ error = &SequenceError{}

type SequenceError struct {
	Sequence *Sequence // The sequence that raised this issue
	Err      error
}

func NewSequenceError(sequence *Sequence, err error) *SequenceError {
	return &SequenceError{
		Sequence: sequence,
		Err:      err,
	}
}

func (s *SequenceError) Error() string {
	return fmt.Errorf("error caused by sequence: %w\n\tSequence: %s", s.Err, s.Sequence.Description).Error()
}

func (s *SequenceError) Is(err error) bool {
	return errors.Is(s.Err, err)
}

func (s *SequenceError) Unwrap() error {
	return s.Err
}
