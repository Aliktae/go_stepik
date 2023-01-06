package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if c < a+b && c > a-b {
		if c*c == a*a+b*b {
			fmt.Println("Прямоугольный")
		} else {
			fmt.Println("Непрямоугольный")
		}
	} else if a < c+b && a > c-b {
		if a*a == c*c+b*b {
			fmt.Println("Прямоугольный")
		} else {
			fmt.Println("Непрямоугольный")
		}
	} else if b < a+c && b > a-c {
		if b*b == a*a+c*c {
			fmt.Println("Прямоугольный")
		} else {
			fmt.Println("Непрямоугольный")
		}
	}
}
