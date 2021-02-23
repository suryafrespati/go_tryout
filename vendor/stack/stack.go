package stack

import "fmt"
import "strings"

func Init() {
	fmt.Println("init Stack")


    names := []string{
        "Surya",
        "Respati",
    }

    var nameAsString string = strings.Join(names, ", ")

    fmt.Println(nameAsString)
}
