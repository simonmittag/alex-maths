package main

import "fmt"

func main() {
	timesFour()
	fmt.Println("")
}

func trueOrFalse() {
	o := 7
	printb(o-2 == 5)
	printb(o*3 == 15)
}

func timesFour() {
	o := []int{2, 9}
	for _, on := range o {
		print(on * 4)
	}
}

func print(o int) {
	fmt.Printf("\n result: %v", o)
}

func printb(o bool) {
	fmt.Printf("\n result: %v", o)
}
