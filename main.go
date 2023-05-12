package main

import "fmt"

func main() {
	timesFour()
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
