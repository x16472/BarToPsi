package main

import (
	"fmt"
)

func main() {
	var (
		bar    float64
		pounds float64
	)
	fmt.Printf("please enter pressure value(bar):")
	fmt.Scanln(&bar)
	pounds = bar * 14.503773773
	fmt.Println(bar, "bar =", pounds, "psi")

}
