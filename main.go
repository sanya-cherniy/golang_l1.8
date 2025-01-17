package main

import "fmt"

func main() {
	test := 256
	bitAssign(&test, false, 8)
	fmt.Println(test)
	bitAssign(&test, true, 8)
	fmt.Println(test)
}

func bitAssign(value *int, bit bool, bitNumber int) {
	if bit {
		*value |= 1 << bitNumber
	} else {
		*value &^= 1 << bitNumber
	}
}
