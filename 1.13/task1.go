package main

import "fmt"

func main() {
	var (
		a           int
		firstDigit  int
		secondDigit int
		thirdDigit  int
		res         int
	)
	fmt.Scan(&a)
	firstDigit = (a / 100) % 10
	secondDigit = (a / 10) % 10
	thirdDigit = a % 10
	res = firstDigit + secondDigit + thirdDigit
	fmt.Println(res)
}
