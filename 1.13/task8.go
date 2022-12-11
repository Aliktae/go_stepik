//Найдите количество минимальных элементов в последовательности.
//Входные данные
//Вводится натуральное число N, а затем N целых чисел последовательности.

//Выходные данные
//Выведите количество минимальных элементов последовательности.

// Sample Input:
// 3
// 21
// 11
// 4
// Sample Output:
// 1
package main

import "fmt"

func main() {
	var n, counter, min int
	fmt.Scan(&n)
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}
	min = numbers[0]
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < min {
			min = numbers[i]
			counter = 1
		} else if numbers[i] == min {
			counter++
		}
	}
	fmt.Println(counter)
}
