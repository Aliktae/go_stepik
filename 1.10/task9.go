//Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.
//
//Входные данные
//Программа получает на вход два числа. Гарантируется, что цифры в числах не повторяются.
//	Числа в пределах от 0 до 10000.
//
//Выходные данные
//Программа должна вывести цифры, которые имеются в обоих числах, через пробел.
//	Цифры выводятся в порядке их нахождения в первом числе.
//
//Sample Input:
//564 8954
//
//Sample Output:
//5 4

package main

import "fmt"

func main() {
	var firstNum, secondNum int
	var firstDigitFirst, secondDigitFirst, thirdDigitFirst, fourthDigitFirst int
	var firstDigitSecond, secondDigitSecond, thirdDigitSecond, fourthDigitSecond int
	fmt.Scan(&firstNum, &secondNum)

	switch {
	case firstNum > 0 && firstNum <= 9:
		firstDigitFirst = firstNum
	case firstNum >= 10 && firstNum <= 99:
		firstDigitFirst = (firstNum / 10) % 10
		secondDigitFirst = firstNum % 10
	case firstNum >= 100 && firstNum <= 999:
		firstDigitFirst = (firstNum / 100) % 10
		secondDigitFirst = (firstNum / 10) % 10
		thirdDigitFirst = firstNum % 10
	case firstNum >= 1000 && firstNum <= 9999:
		firstDigitFirst = (firstNum / 1000) % 10
		secondDigitFirst = (firstNum / 100) % 10
		thirdDigitFirst = (firstNum / 10) % 10
		fourthDigitFirst = firstNum % 10
	}

	switch {
	case secondNum > 0 && secondNum <= 9:
		firstDigitSecond = secondNum
	case secondNum >= 10 && secondNum <= 99:
		firstDigitSecond = (secondNum / 10) % 10
		secondDigitSecond = secondNum % 10
	case secondNum >= 100 && secondNum <= 999:
		firstDigitSecond = (secondNum / 100) % 10
		secondDigitSecond = (secondNum / 10) % 10
		thirdDigitSecond = secondNum % 10
	case secondNum >= 1000 && secondNum <= 9999:
		firstDigitSecond = (secondNum / 1000) % 10
		secondDigitSecond = (secondNum / 100) % 10
		thirdDigitSecond = (secondNum / 10) % 10
		fourthDigitSecond = secondNum % 10
	}

	if firstDigitFirst == firstDigitSecond ||
		firstDigitFirst == secondDigitSecond ||
		firstDigitFirst == thirdDigitSecond ||
		firstDigitFirst == fourthDigitSecond {
		fmt.Print(firstDigitFirst, " ")
	}

	if secondDigitFirst == firstDigitSecond ||
		secondDigitFirst == secondDigitSecond ||
		secondDigitFirst == thirdDigitSecond ||
		secondDigitFirst == fourthDigitSecond {
		fmt.Print(secondDigitFirst, " ")
	}

	if thirdDigitFirst == firstDigitSecond ||
		thirdDigitFirst == secondDigitSecond ||
		thirdDigitFirst == thirdDigitSecond ||
		thirdDigitFirst == fourthDigitSecond {
		fmt.Print(thirdDigitFirst, " ")
	}

	if fourthDigitFirst == firstDigitSecond ||
		fourthDigitFirst == secondDigitSecond ||
		fourthDigitFirst == thirdDigitSecond ||
		fourthDigitFirst == fourthDigitSecond {
		fmt.Print(fourthDigitFirst, " ")
	}

}
