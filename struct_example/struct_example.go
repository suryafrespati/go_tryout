package struct_example

import "fmt"

type person struct {
    name string
    age  int
}

type student struct {
    name string
    grade int
}

func Init() {
    fmt.Println("Init struct_example")

    var s1 student
    s1.name = "john wick"
    s1.grade = 2

    fmt.Println("name  :", s1.name)
    fmt.Println("grade :", s1.grade)

    initializationExample()
    sliceAndStructExample()
    anonymousStructExample()
}

func initializationExample() {
    var s1 = student{}
    s1.name = "wick"
    s1.grade = 2

    var s2 = student{"ethan", 2}
    var s3 = student{name: "jason"}

    fmt.Println("student 1 :", s1.name)
    fmt.Println("student 2 :", s2.name)
    fmt.Println("student 3 :", s3.name)

    // var s4 = student{name: "wayne", grade: 2}
    // var s5 = student{grade: 2, name: "bruce"}
}

func sliceAndStructExample() {
    fmt.Println("sliceAndStructExample()")

    var allStudents = []person{
        {name: "Wick", age: 23},
        {name: "Ethan", age: 23},
        {name: "Bourne", age: 22},
    }

    // for _, student := range allStudents {
    //     fmt.Println(student.name, "age is", student.age)
    // }

    for i := 0; i < len(allStudents); i++ {
        student := allStudents[i]
        fmt.Println(student.name, "age is", student.age)
    }
}

func anonymousStructExample() {
    fmt.Println("anonymousStructExample()")

    var s1 = struct {
        person
        grade int
    }{}
    s1.person = person{"martin", 27}
    s1.grade = 10

    fmt.Println("name  :", s1.person.name)
    fmt.Println("age   :", s1.person.age)
    fmt.Println("grade :", s1.grade)
}

