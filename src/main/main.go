package main

import (
	"fmt"
	"strings"
)

func main() {
	simpleChannel := make(chan string, 1)
	simpleChannel <- "Good Job!"
	fmt.Println(<-simpleChannel)

	fmt.Println("==================================")
	fmt.Println("")

	message := "Chuck Norris doesn't jump. He moves the ground away from him."

	parts := strings.Split(message, " ")
	chucksChannel := make(chan string, len(parts))

	for _, word := range parts {
		chucksChannel <- word
	}

	//	close(chucksChannel) // only entrance closed

	for i := 0; i < len(parts); i++ {
		fmt.Print(<-chucksChannel, " ")
	}

	//	for {
	//		if str, pullResult := <-chucksChannel; pullResult {
	//			fmt.Print(str, " ")
	//
	//		} else {
	//			break
	//		}
	//	}

	//	chucksChannel <- "you will get an error if you send message after close"

	fmt.Println("==================================")
	fmt.Println("")
}
