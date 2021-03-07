package embedded_struct_example

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	semester byte
	major
	person
}

type major struct {
	name   string
	degree string
}

func Init() {
	fmt.Printf("\n# --- embedded_struct_example\n\n")

	var s1 = student{}
	s1.person.name = "Aldo Raine"
	s1.person.age = 22
	s1.major.name = "Statistic"
	s1.major.degree = "Bachelor"
	s1.semester = 7

	fmt.Println("name  :", s1.person.name)
	fmt.Println("age   :", s1.person.age)
	fmt.Println("major :", s1.major.name)
	fmt.Println("degree :", s1.major.degree)
	fmt.Println("semester :", s1.semester)

	fmt.Printf("\n# --- end of embedded_struct_example\n\n")
}
