package error_example

import (
	"errors"
	"fmt"
	"strings"
)

func Init() {
	fmt.Println("Init error_example")

	var name string = ""
	// var name string = "Bruce Wayne"

	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		fmt.Println(err.Error())
	}

	fmt.Println()
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}
