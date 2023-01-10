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

	for i := 1; i < 10; i++ {
		fmt.Println(i * i)
	}

    var users = [3]string{"Tom", "Alice", "Kate"}
    for index, value := range users
    {
        fmt.Println(index, value)
    }
}