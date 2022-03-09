package main

import (
	"fmt"
	"time"
)


func timer(runner func()) {
	start := time.Now()
	runner()
	fmt.Println()
	fmt.Printf("Function took %f seconds\n", time.Since(start).Seconds())
}


func generateNumbers(end int) []int {
	numbers := []int{}
	for i := 1; i <= end; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}

func main() {
	timer(func() {
		fmt.Printf("%v\n", generateNumbers(1_000_000))
	})
}