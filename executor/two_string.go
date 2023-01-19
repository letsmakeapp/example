package executor

import "example/abstract"

type TwoStringLengthCounter struct {
	Counter abstract.StringLengthCounter
}

func (c TwoStringLengthCounter) Count(a, b string) uint {
	return c.Counter.Count(a) + c.Counter.Count(b)
}
