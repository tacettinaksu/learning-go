package main

import "fmt"

func main() {

	foo := 3

	switch {
	case foo == 1:
		fmt.Println("one")
	case foo > 2:
		fmt.Println(">2")
	}

	if foo == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("other than one")
	}

	fmt.Println("")
	fmt.Println("-----------------------------------------------")

	m := make(map[string]string)

	m["java"] = "good"
	m[".net"] = "bad"
	m["go"] = "cool"

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("")
	fmt.Println("-----------------------------------------------")

	arr := []string{"scrum", "agile", "lean", "xp"}

	for idx, value := range arr {
		fmt.Println(idx, value)
	}

	fmt.Println("")
	fmt.Println("-----------------------------------------------")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("")
	fmt.Println("-----------------------------------------------")
	j := 0

	for {
		j++
		fmt.Println(j)

		if j > 5 {
			break
		}
	}
}
