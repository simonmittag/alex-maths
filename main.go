package main

import "fmt"

func main() {
	timesX(7)
	fmt.Println("")
}

func trueOrFalse() {
	o := 7
	print(o-2 == 5)
	print(o*3 == 15)
}

func timesX(x int) {
	fmt.Printf("\n X is %v", x)
	o := []int{2, 9}
	for _, on := range o {
		print(on * x)
	}
}

func print(o interface{}) {
	fmt.Printf("\n result: %v", o)
}
