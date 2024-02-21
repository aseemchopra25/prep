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

// 2.

func atopowerb(a int, b int) int {
	if b < 0 {
		return 0
	} else if b == 0 {
		return 1
	} else {
		return a * atopowerb(a, b-1)
	}

}

// 3.

func mod(a int, b int) int {
	if b <= 0 {
		return -1
	}
	div := a / b
	return a - div*b
}

// 4.
func div(a int, b int) int {
	count := 0
	sum := b
	for sum <= a {
		sum += b
		count++
	}
	return count

}

func main() {
	fmt.Println(product(10, 20)) // Big-O: O(b)
	fmt.Println(atopowerb(5, 2)) // Big-O: O(b)
	fmt.Println(mod(10, 3))      // Big-O: O(1)
	fmt.Println(div(10, 3))      // Big-O: O(a/b)

}
