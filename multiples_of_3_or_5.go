// https://www.codewars.com/kata/514b92a657cdc65150000006
package main

import "fmt"

func isDevidedBy(src, devider int) bool {
	return src%devider == 0
}

func Multiple3And5(number int) int {
	sum := 0

	for i := 1; i < number; i++ {
		if isDevidedBy(i, 3) || isDevidedBy(i, 5) {
			sum += i
		}
	}

	return sum
}

func main() {
	res := Multiple3And5(10) // 23
	fmt.Printf("Res: %d\n", res)
}
