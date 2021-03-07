package main

import (
	"fmt"
	"tryout/array_slice_example"
	"tryout/embedded_struct_example"
	"tryout/error_example"
	"tryout/func_example"
	"tryout/goroutine_example"
	"tryout/interface_example"
	"tryout/map_example"
	"tryout/method_example"
	"tryout/pointer_example"
	"tryout/reflect_example"
	"tryout/stack"
	"tryout/struct_example"
)

func main() {
	fmt.Printf("\n@@@ learn Golang\n\n")

	stack.Init()
	array_slice_example.Init()
	map_example.Init()
	func_example.Init()
	pointer_example.Init()
	struct_example.Init()
	embedded_struct_example.Init()
	method_example.Init()
	interface_example.Init()
	reflect_example.Init()
	error_example.Init()
	goroutine_example.Init()

	fmt.Printf("\n@@@ end of learn Golang\n\n")
}
