package main

import (
	"fmt"
	"time"
)

func main() {
	//go routine
	go func() {
		for i := 0; i < 150; i++ {
			fmt.Println("Go Johnny Go")
		}
	}()

	go func() {
		for i := 0; i < 150; i++ {
			fmt.Println("Back to future!")
		}
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}
