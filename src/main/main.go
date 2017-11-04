package main

import "fmt"

func main() {
	var myInt int
	myInt = 42

	var myFloat32 float32 = 42.

	myString := "Hello Go"

	myComplex := complex(2, 3)

	println(myInt)
	println(myFloat32)
	println(myString)
	println(myComplex)
	println(real(myComplex))
	println(imag(myComplex))

	myArray := [...]int{28, 53, 61, 52, 29}

	fmt.Println("")
	fmt.Println("------------ ARRAYS -----------------------------")

	myArray[2] = 19
	//myArray[5] = 23     //-> ERROR on compile time

	fmt.Println("myArray ", myArray)

	mySlice := myArray[:]

	fmt.Println("mySlice full      : ", mySlice)

	mySlice = myArray[2:4]

	fmt.Println("mySlice 2:4       : ", mySlice)

	mySlice = append(mySlice, 100)

	fmt.Println("mySlice appended  : ", mySlice)

	fmt.Println("")
	fmt.Println("---------------   MAPS   ------------------------------------")

	myMap := make(map[int]string)

	fmt.Println("myMap empty  : ", myMap)

	myMap[42] = "Foo"
	myMap[12] = "Bar"

	fmt.Println("myMap        : ", myMap)
	fmt.Println("myMap key 12 : ", myMap[12])
}
