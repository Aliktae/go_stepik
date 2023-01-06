package main

import "fmt"

func reverseNumber(a int) int {
	res := 0
	for a > 0 {
		r := a % 10
		res = (res * 10) + r
		a /= 10
	}
	return res
}

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(reverseNumber(a))
}
