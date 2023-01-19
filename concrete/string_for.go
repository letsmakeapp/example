package concrete

import "example/abstract"

type ForLoopStringLengthCounter struct{}

var (
	_ abstract.StringLengthCounter = (*ForLoopStringLengthCounter)(nil)
)

func (f ForLoopStringLengthCounter) Count(s string) uint {
	length := uint(0)
	for range s {
		length++
	}
	return length
}
