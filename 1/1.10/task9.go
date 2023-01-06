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

	if secondDigitFirst != 0 && secondDigitFirst == firstDigitSecond ||
		secondDigitFirst != 0 && secondDigitFirst == secondDigitSecond ||
		secondDigitFirst != 0 && secondDigitFirst == thirdDigitSecond ||
		secondDigitFirst != 0 && secondDigitFirst == fourthDigitSecond {
		fmt.Print(secondDigitFirst, " ")
	}

	if thirdDigitFirst != 0 && thirdDigitFirst == firstDigitSecond ||
		thirdDigitFirst != 0 && thirdDigitFirst == secondDigitSecond ||
		thirdDigitFirst != 0 && thirdDigitFirst == thirdDigitSecond ||
		thirdDigitFirst != 0 && thirdDigitFirst == fourthDigitSecond {
		fmt.Print(thirdDigitFirst, " ")
	}

	if fourthDigitFirst != 0 && fourthDigitFirst == firstDigitSecond ||
		fourthDigitFirst != 0 && fourthDigitFirst == secondDigitSecond ||
		fourthDigitFirst != 0 && fourthDigitFirst == thirdDigitSecond ||
		fourthDigitFirst != 0 && fourthDigitFirst == fourthDigitSecond {
		fmt.Print(fourthDigitFirst)
	}

}
