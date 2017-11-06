package main

import (
	"runtime"
	"time"
)

func main() {

	//runtime.GOMAXPROCS(2)
	dur10ms, _ := time.ParseDuration("10ms")

	go func() {
		for i := 0; i < 150; i++ {
			println("Go Johnny Go")
			time.Sleep(dur10ms)
		}
	}()

	go func() {
		for i := 0; i < 150; i++ {
			println("Back to future!")
			time.Sleep(dur10ms)
		}
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}
