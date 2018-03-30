package schema

import (
	"fmt"
	"sync"
)

// Sequence threadsafe incrementing sequence of integers.
type Sequence struct {
	Name   string `json:"name"`
	Number uint64 `json:"number"`
	lock   sync.Mutex
}

// NewSequence creates a new sequence.
func NewSequence(name string) *Sequence {
	return &Sequence{
		Name:   fmt.Sprintf("sequence%s%s", KeyDelimiter, name),
		Number: 0,
		lock:   sync.Mutex{},
	}
}

// Next returns the next number in the sequence.
func (s *Sequence) Next() uint64 {
	s.lock.Lock()
	defer s.lock.Unlock()
	next := s.Number + 1
	s.Number = next
	return next
}
