// https://www.codewars.com/kata/521c2db8ddc89b9b7a0000c1/train/go
package main

import (
	"fmt"
)

type Direction int

const (
	Right Direction = iota
	Down
	Left
	Up
)

func Snail(snaipMap [][]int) []int {
	fmt.Printf("%v", snaipMap)

	matrixSize := len(snaipMap) * len(snaipMap)
	res := make([]int, 0, matrixSize)
	if matrixSize == 0 {
		return res
	}

	direction := Right
	pX := 0
	pY := 0
	borderRight := len(snaipMap) - 1
	borderLeft := 0
	borderTop := 0
	borderBottom := len(snaipMap) - 1

	for i := 0; i < matrixSize; i++ {
		res = append(res, snaipMap[pY][pX])
		fmt.Printf("direction: %d\n", direction)

		switch direction {
		case Right:
			if pX >= borderRight {
				direction = Down
				borderTop++
			}
		case Down:
			if pY >= borderBottom {
				direction = Left
				borderRight--
			}
		case Left:
			if pX <= borderLeft {
				direction = Up
				borderBottom--
			}
		case Up:
			if pY <= borderTop {
				direction = Right
				borderLeft++
			}
		}

		switch direction {
		case Right:
			pX++
		case Down:
			pY++
		case Left:
			pX--
		case Up:
			pY--
		}
	}

	return res
}

func main() {
	snailMap := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	res := Snail(snailMap)
	fmt.Printf("res: %v\n", res) // 1, 2, 3, 6, 9, 8, 7, 4, 5
}
