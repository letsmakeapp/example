package main

import (
	"example/concrete"
	"fmt"

	"example/executor"
)

func main() {
	counter := executor.TwoStringLengthCounter{
		Counter: concrete.ForLoopStringLengthCounter{},
	}

	count := counter.Count("HELLO", "WORLD!")
	fmt.Println(count)
}
