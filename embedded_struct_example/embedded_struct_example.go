package embedded_struct_example

import "fmt"

type person struct {
    name string
    age  int
}

type student struct {
    // age int
    grade int
    person
}

func Init() {
    fmt.Println("Init embedded_struct_example")

    var s1 = student{}
    s1.name = "wick"
    s1.age = 21
    s1.person.age = 14
    s1.grade = 2

    fmt.Println("name  :", s1.name)
    fmt.Println("age   :", s1.age)
    fmt.Println("age   :", s1.person.age)
    fmt.Println("grade :", s1.grade)

    subStructExample()
}

func subStructExample() {
    fmt.Println("subStructExample()")

    var p1 = person{name: "wick", age: 21}
    var s1 = student{person: p1, grade: 2}

    fmt.Println("name  :", s1.name)
    fmt.Println("age   :", s1.age)
    fmt.Println("grade :", s1.grade)
}

