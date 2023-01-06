//Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является, то есть
//выведите такое число n, что φn=A. Если А не является числом Фибоначчи, выведите число -1.
//
//Входные данные
//Вводится натуральное число.
//
//Выходные данные
//Выведите ответ на задачу.
//
//Sample Input:
//
//8
//Sample Output:
//
//6

package main

import "fmt"

func main() {
	var n, b int
	fmt.Scan(&n)
	b = 0
	var fib []int = []int{0, 1}
	for i := 1; i < n; i++ {
		fib = append(fib, fib[i]+fib[i-1])
		if n == 2 {
			fmt.Println(3)
		} else if n == fib[i] {
			fmt.Println(i)
			b = 0
			break
		} else {
			b = -1
		}
	}
	if b == -1 {
		fmt.Println(b)
	}
}
