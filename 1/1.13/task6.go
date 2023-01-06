package main

import "fmt"

func main() {
	var a, b float32
	fmt.Scan(&a, &b)
	var res float32
	res = (a + b) / 2
	fmt.Println(res)
}
