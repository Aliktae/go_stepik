package main

import "fmt"

func main() {
	var arr int
	for fmt.Scan(&arr); arr <= 100; fmt.Scan(&arr) {
		if arr >= 10 && arr <= 100 {
			fmt.Println(arr)
			continue
		}
	}
}
