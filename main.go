package main

import "fmt"

func main() {
	timesFour()
	fmt.Println("")
}

func trueOrFalse() {
	o := 7
	print(o-2 == 5)
	print(o*3 == 15)
}

func timesFour() {
	o := []int{2, 9}
	for _, on := range o {
		print(on * 4)
	}
}

func print(o interface{}) {
	fmt.Printf("\n result: %v", o)
}
