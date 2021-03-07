package error_example

import (
	"errors"
	"fmt"
	"strings"
)

func Init() {
	fmt.Printf("\n# --- error_example\n\n")

	var name string = ""
	// var name string = "Bruce Wayne"

	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		fmt.Println(err.Error())
	}

	fmt.Printf("\n# --- end of error_example\n\n")
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("input cannot be empty")
	}
	return true, nil
}
