package main

import "fmt"

func main() {
	var time int
	fmt.Scan(&time)
	hours := time / 3600
	mins := (time % 3600) / 60
	fmt.Println("It is", hours, "hours", mins, "minutes.")
}
