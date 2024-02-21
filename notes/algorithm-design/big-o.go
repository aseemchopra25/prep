package main

import "fmt"

// 1.
func product(a int, b int) int {
	sum := 0
	for i := 0; i < b; i++ {
		sum += a
	}
	return sum
}

func main() {
	fmt.Println(product(10, 20)) // Big-O: O(b)
}
