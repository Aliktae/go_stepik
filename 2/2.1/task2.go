//Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.
//
//Входные данные
//
//Вводится четыре числа.
//
//Выходные данные
//
//Необходимо вернуть из функции наименьшее из 4-х данных чисел
//
//Sample Input:
//
//4 5 6 7
//Sample Output:
//
//4
//package main
//
//func minimumFromFour() int {
//	//print your code here
//}

package main

import "fmt"

func minimumFromFour() int {
	var (
		x, y, z, d int
	)
	fmt.Scan(&x, &y, &z, &d)
	min := 0
	switch {
	case x <= y && x <= z && x <= d:
		min = x
	case y <= x && y <= z && y <= d:
		min = y
	case z <= x && z <= y && z <= d:
		min = z
	case d <= x && d <= y && d <= z:
		min = d
	}
	return min
}
