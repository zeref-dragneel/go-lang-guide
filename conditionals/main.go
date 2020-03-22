package main

import "fmt"

func main() {
	x := 3
	y := 3
	//simple if
	if x%2 == 0 {
		fmt.Println("x is even")
	}
	//if -else
	if x%2 == 0 {
		fmt.Println("x is even")
	} else {
		fmt.Println("x is odd")
	}
	//if else if else
	if x < y {
		fmt.Println("x is greater")
	} else if y < x {
		fmt.Println("x is odd")
	} else {
		fmt.Println("equal")
	}
	color := "redi"
	switch color {
	case "red":
		fmt.Println("stop")
	default:
		fmt.Println("go")
	}

}
