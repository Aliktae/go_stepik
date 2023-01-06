package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	var counter int
	for i := 0; i < n; i++ {
		if arr[i] >= 0 {
			counter++
		}
	}
	fmt.Println(counter)

}
