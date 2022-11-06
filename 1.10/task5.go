package main

import "fmt"

func main() {
	var (
		max     int = 0
		a       int
		counter int = 0
	)
	for fmt.Scan(&a); a != 0; fmt.Scan(&a) {
		if max < a {
			max = a
			counter = 0
		}
		if max == a {
			counter += 1
		}
	}
	fmt.Println(counter)
}
