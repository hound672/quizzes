// https://www.codewars.com/kata/5526fc09a1bbd946250002dc/train/go
package main

import (
	"fmt"
)

func FindOutlier(integers []int) int {
	var odds []int
	var evens []int

	for _, el := range integers {
		if el%2 == 0 {
			evens = append(evens, el)
		} else {
			odds = append(odds, el)
		}
	}

	if len(odds) == 1 {
		return odds[0]
	} else {
		return evens[0]
	}
}

func main() {
	//res := FindOutlier([]int{2, 6, 8, -10, 3})
	//res := FindOutlier([]int{math.MaxInt32, 0, 1})
	res := FindOutlier([]int{206847684, 1056521, 7, 17, 1901, 21104421, 7, 1, 35521, 1, 7781})

	fmt.Printf("Res: %v\n", res) // 3
}
