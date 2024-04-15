package main

import (
	"fmt"
)

func main() {
	var a, b int = 2, 3
	message := "Hello world!"

	numbers := make([]int, 5, 10)

	numbers = append(numbers, 1, 2)

	fmt.Println(message)
	fmt.Println("sum: ", a+b)
	fmt.Println("numbers: ", numbers)
	// fmt.Println(os.Getenv("APP_NAME"))
}
