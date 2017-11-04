package main

import (
	"fmt"
)

func main() {

	studentJohn := Student{}
	studentJohn.name = "John Doe"
	studentJohn.number = 41
	fmt.Println(studentJohn)

	fmt.Println("")
	fmt.Println("==============================================")

	studentRob := Student{"Rob J", 65}
	fmt.Println(studentRob)

	fmt.Println("")
	fmt.Println("==============================================")

	studentJane := new(Student)
	studentJane.name = "Jane Z"
	studentJane.number = 312

	fmt.Println(*studentJane)
	fmt.Println(studentJane)
	println("===", studentJane)

	fmt.Println("")
	fmt.Println("==============================================")

	studentMary := &Student{"Mary Jane", 566}
	fmt.Println(studentMary)

	fmt.Println("")
	fmt.Println("==============================================")

	studentPeter := &Student{"Peter Pan", 566}
	fmt.Println(studentPeter)

	fmt.Println("")
	fmt.Println("==============================================")

	studentJames := &Student{"James 7", 7}
	fmt.Println(studentJames)
	//add function to struct. you can add new method to structs from anywhere
	fmt.Println(studentJames.isNumberEven())
}

type Student struct {
	name   string
	number int
}

func newStudent(name string, number int) *Student {
	result := Student{"name", number}
	return &result
}

func (s *Student) isNumberEven() bool {
	if s.number%2 == 0 {
		return true
	}
	return false
}
