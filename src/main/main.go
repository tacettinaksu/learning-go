package main

import "fmt"

func main() {
	message := "Hello"
	sayHello_callByValue(message)
	fmt.Println(message)

	sayHello_callByReferance(&message)
	fmt.Println(message)

	fmt.Println("======================================")

	sayHelloMultiple("Talk", "is", "cheap", "show", "me", "the", "code")

	fmt.Println("============= MULTIPLE RETURN ====================")

	divison, remainder := divideOperation(11, 3)
	fmt.Println("11/3=", divison, " Remainder= ", remainder)

	fmt.Println("")
	fmt.Println("============Anonymous function=====================")

	powerOperation := func(base int, power int) (result int) {
		result = base
		for i := 1; i < power; i++ {
			result *= base
		}
		return result
	}

	fmt.Println("2^7= ", powerOperation(2, 7))
}

func sayHello_callByValue(message string) {
	fmt.Println(message)
	message = "Hello updated by callByValue."
}

func sayHello_callByReferance(message *string) {
	fmt.Println(*message)
	*message = "Hello updated by callByReferance."
}

func sayHelloMultiple(messages ...string) {
	for _, message := range messages {
		fmt.Println(message)
	}
}

func divideOperation(dividend int, divisor int) (divison int, remainder int) {
	divison = dividend / divisor
	remainder = dividend - (divison * divisor)
	return
}
