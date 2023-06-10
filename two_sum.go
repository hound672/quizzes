// https://www.codewars.com/kata/52c31f8e6605bcc646000082
package main

import "fmt"

func TwoSum(numbers []int, target int) [2]int {
	seen := make(map[int]int)

	for i, number := range numbers {
		match := target - number

		if j, found := seen[match]; found {
			return [2]int{j, i}
		}

		seen[number] = i
	}

	return [2]int{}
}

func main() {
	res := TwoSum([]int{1, 2, 3}, 4) // returns [2]int{0, 2}
	fmt.Printf("res: %v\n", res)
}
