package main

import "fmt"

func main() {
	// var msg string
	// msg = "Hello Go!"

	// age := "25"

	// var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	// fmt.Println(msg + age)

	numbers := [...]int{1, 2, 3, 4, 5}

	fmt.Println(numbers)
	fmt.Println(numbers[2])
}
