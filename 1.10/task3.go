package main

import "fmt"

func main() {
	var (
		a   int
		b   int
		sum int
	)
	sum = 0
	fmt.Scan(&a, &b)
	for ; a <= b; a++ {
		sum += a
	}
	fmt.Println(sum)
}
