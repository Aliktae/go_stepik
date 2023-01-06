package main

import "fmt"

func main() {
	var x, p, y int
	fmt.Scan(&x, &p, &y)
	year := 0
	for {
		if x > y {
			break
		} else {
			x += x * p / 100
			year += 1
		}
	}
	fmt.Println(year)
}
