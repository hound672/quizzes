// https://www.codewars.com/kata/52597aa56021e91c93000cb0/train/go
package main

import "fmt"

func MoveZeros(arr []int) []int {
	res := make([]int, len(arr))
	pos := 0

	for i := range arr {
		if arr[i] != 0 {
			res[pos] = arr[i]
			pos++
		}
	}

	return res
}

func main() {
	res := MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}) // To(Equal([]int{ 1, 2, 1, 1, 3, 1, 0, 0, 0, 0 })
	fmt.Printf("res: %v\n", res)
}
