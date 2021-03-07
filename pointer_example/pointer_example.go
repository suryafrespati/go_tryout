package pointer_example

import "fmt"

func Init() {
	fmt.Printf("\n# --- pointer_example\n\n")

	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)

	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address) :", numberB)

	var number = 4
	fmt.Println("before :", number)

	change(&number, 10)
	fmt.Println("after  :", number)

	fmt.Printf("\n# --- end of pointer_example\n\n")
}

func change(original *int, value int) {
	*original = value
}
