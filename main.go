package main

import (
	"fmt"
	"math"
)

func calc(number int, total *int) int {
	if number == 0 {
		return *total
	}
	*total += int(math.Pow(float64(number%10), float64(2)))
	return calc(number/10, total)
}

func is_happy_number(number int) bool {
	var i int
	for i = 100; i > 0; i-- {
		total := 0
		number = calc(number, &total)
		if number == 1 {
			return true
		}
	}
	return false
}

func main() {
	for n := 1; n <= 1000; n += 1 {
		if is_happy_number(n) {
			fmt.Printf("%d, ", n)
		}
	}
}
