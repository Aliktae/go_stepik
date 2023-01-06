//По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в
//порядке возрастания.
//
//Входные данные
//Вводится натуральное число.
//
//Выходные данные
//Выведите ответ на задачу.
//
//Sample Input:
//
//50
//Sample Output:
//
//1 2 4 8 16 32

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	base := 1
	for i := 0; i < n; i++ {
		if base > n {
			break
		}
		fmt.Print(base, " ")
		base *= 2
	}
}
