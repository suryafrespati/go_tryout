package main

import (
	"array_slice_example"
	"fmt"
	"func_example"
	"map_example"
	"stack"
)

func main() {
	fmt.Println("Hello, World!")
	print("\n\n\n")

	stack.Init()
	array_slice_example.Init()
	map_example.Init()
	func_example.Init()

    fmt.Println("--- end of file")
}
