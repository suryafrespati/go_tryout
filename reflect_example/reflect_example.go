package reflect_example

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func Init() {
	fmt.Println("Init reflect_example")

	var number = 23
	var reflectValue = reflect.ValueOf(number)

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel :", reflectValue.Int())
	}

	var s1 = &student{Name: "adam", Grade: 9}
	s1.getPropertyInfo()
}

func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama      :", reflectType.Field(i).Name)
		fmt.Println("tipe data :", reflectType.Field(i).Type)
		fmt.Println("nilai     :", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}
