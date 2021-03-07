package stack

import (
	"fmt"
)

func Init() {
	fmt.Printf("\n# --- stack\n\n")

	var nums []int = []int{
		2, 4, 6, 8, 12, 24, 100, 128,
	}

	var stack = NewStack()

	for _, v := range nums {
		stack.Push(v)
	}

	stack.Show()

	stack.Pop()
	stack.Pop()
	stack.Pop()

	stack.Show()

	var stack2 Stack
	stack2.Show()

	fmt.Printf("\n# --- end of stack\n\n")
}
