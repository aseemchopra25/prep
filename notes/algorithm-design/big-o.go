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

func atopowerb(a int, b int) int {
	if b < 0 {
		return 0
	} else if b == 0 {
		return 1
	} else {
		return a * atopowerb(a, b-1)
	}

}

func main() {
	fmt.Println(product(10, 20)) // Big-O: O(b)
	fmt.Println(atopowerb(5, 2)) // Big-O: O(b)
}
