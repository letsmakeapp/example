package concrete

import "example/abstract"

type NativeStringLengthCounter struct{}

var (
	_ abstract.StringLengthCounter = (*NativeStringLengthCounter)(nil)
)

func (n NativeStringLengthCounter) Count(s string) uint {
	return uint(len(s))
}
