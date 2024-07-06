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

func MoveZerosN2(arr []int) []int {
	k := 1

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			for j := i + k; j < len(arr); j++ {
				if arr[j] == 0 {
					continue
				}
				arr[i] = arr[j]
				arr[j] = 0
				i++
				k = j - i
			}
		}
	}

	return arr
}

func MoveZerosN(arr []int) []int {
	k := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[k] = arr[i]
			k++
			continue
		}
	}
	for i := k; i < len(arr); i++ {
		arr[i] = 0
	}

	return arr
}

func main() {
	example := []int{1, 2, 0, 0, 0, 0, 2, 0, 1, 0, 3, 0, 1, 111, 0, 222}

	arrMem := make([]int, len(example))
	copy(arrMem, example)
	resMem := MoveZerosMem(arrMem)
	fmt.Printf("res: %v\n", resMem)

	arrNN := make([]int, len(example))
	copy(arrNN, example)
	resNN := MoveZerosMem(arrNN)
	fmt.Printf("res: %v\n", resNN)

	arrN := make([]int, len(example))
	copy(arrN, example)
	resN := MoveZerosN(arrN)
	fmt.Printf("res: %v\n", resN)
}
