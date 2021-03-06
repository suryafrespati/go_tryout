package stack

import (
	"fmt"
	"strings"
)

func Init() {
	fmt.Println("init Stack")

	names := []string{
		"Surya",
		"Respati",
	}

	var nameAsString string = strings.Join(names, ", ")

	fmt.Println(nameAsString)
}
