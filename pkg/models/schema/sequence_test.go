package schema

import (
	"testing"
)

func TestSequenceNext(t *testing.T) {
	seq := NewSequence("test")
	nextNum := seq.Next()
	if 1 != nextNum {
		t.Errorf("sequence.Next returned wrong number: Expected=1 Got=%d", nextNum)
	}
}
