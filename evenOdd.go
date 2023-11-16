package main

import "fmt"

func main() {
	var sum int = 0

	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum += i
		} else {
			sum -= i
		}
	}

	fmt.Println("Sum value:", sum)
}
