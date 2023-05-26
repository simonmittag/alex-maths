package main

import "fmt"

func main() {
	timesX(4)
	fmt.Println("")
}

func trueOrFalse() {
	o := 7
	print(o-2 == 5)
	print(o*3 == 15)
}

func timesX(x int) {
	o := []int{2, 9}
	for _, on := range o {
		print(on * x)
	}
}

func print(o interface{}) {
	fmt.Printf("\n result: %v", o)
}
