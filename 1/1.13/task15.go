package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	result := 0
	for i := 1; a > 0; {
		if a%10 != b {
			result += a % 10 * i
			i *= 10
		}
		a /= 10
	}
	fmt.Println(result)
}
