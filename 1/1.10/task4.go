package main

import "fmt"

func main() {
	var size int
	fmt.Scanln(&size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	sum := 0
	for i := 0; i < size; i++ {
		if arr[i] > 10 && arr[i] < 100 && arr[i]%8 == 0 {
			sum += arr[i]
		}
	}
	fmt.Println(sum)
}
