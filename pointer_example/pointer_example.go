package pointer_example

import "fmt"

func Init() {
    fmt.Println("Init pointer_example")

    var numberA int = 4
    var numberB *int = &numberA

    fmt.Println("numberA (value)   :", numberA)  // 4
    fmt.Println("numberA (address) :", &numberA) // 0xc20800a220

    fmt.Println("numberB (value)   :", *numberB) // 4
    fmt.Println("numberB (address) :", numberB)  // 0xc20800a220


    var number = 4
    fmt.Println("before :", number) // 4

    change(&number, 10)
    fmt.Println("after  :", number) // 10

    var daysOfYear int = 365
    fmt.Println("daysOfYear", daysOfYear)

    var durationOfStudy *int = &daysOfYear
    fmt.Println("durationOfStudy", *durationOfStudy)
}

func change(original *int, value int) {
    *original = value
}

