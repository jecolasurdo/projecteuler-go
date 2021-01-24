package main

import (
	"fmt"
	"strconv"

	"github.com/jecolasurdo/projecteuler/p705"
)

func main() {
	for i := 1111; i < 10000; i++ {
		s := strconv.Itoa(i)
		swaps := p705.InversionCount([]byte(s))
		// fmt.Println(s, swaps)
		fmt.Println(swaps)
	}
}
