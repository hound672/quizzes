// https://www.codewars.com/kata/526571aae218b8ee490006f4
package main

import "fmt"

func CountBits(src uint) int {
	offset := 0
	count := 0

	for src >= (1 << offset) {
		if (src & (1 << offset)) > 0 {
			count++
		}
		offset++
	}
	return count
}

func main() {
	res := CountBits(1234)
	fmt.Printf("Res: %d\n", res)
}
