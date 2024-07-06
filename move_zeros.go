// https://www.codewars.com/kata/52597aa56021e91c93000cb0/train/go
package main

import "fmt"

func MoveZerosMem(arr []int) []int {
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

func MoveZeros(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			for j := i + 1; j < len(arr); j++ {
				if arr[j] == 0 {
					continue
				}
				arr[i] = arr[j]
				arr[j] = 0
				//i++
				i = j + 1
			}
		}
	}

	return arr
}

func main() {
	res := MoveZeros([]int{1, 2, 0, 1, 0, 0, 2, 0, 1, 0, 3, 0, 1}) // To(Equal([]int{ 1, 2, 1, 1, 3, 1, 0, 0, 0, 0 })
	fmt.Printf("res: %v\n", res)
}
