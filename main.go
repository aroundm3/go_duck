package main

import (
	"fmt"
)

func main() {

	//data types
	var a, b int = 2, 3
	message := "Hello world!"

	testArr := [0]int{}
	testArr2 := [3]int{1, 2}

	numbers := make([]int, 5, 10)
	fmt.Println("numbers: ", numbers)

	numbers = append(numbers, 1, 2)
	numbers2 := make([]int, 12)

	copy(numbers2, numbers)

	testSubSlice := []int{2, 4, 1, 6, 9, 0}

	subSlice1 := testSubSlice[2:] //[1 6 9 0]
	subSlice2 := testSubSlice[:2] //[2 4]

	fmt.Println(message)
	fmt.Println("sum: ", a+b)
	fmt.Println("numbers: ", numbers)
	fmt.Println("numbers2: ", numbers2)
	fmt.Println("testArr", testArr)
	fmt.Println("testArr2", testArr2)
	fmt.Println("subSlice1", subSlice1)
	fmt.Println("subSlice2", subSlice2)

	//loop
	i := 0
	sum := 0
	for i < 10 {
		sum += 1
		i++
	}
	fmt.Println(sum)

	//pointer
	var ap *int
	demo := 12
	ap = &demo

	fmt.Println("pointer: ", ap)
	// fmt.Println(os.Getenv("APP_NAME"))
}
