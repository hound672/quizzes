// https://www.codewars.com/kata/525c65e51bf619685c000059/train/go
package main

import "log"

func Cakes(recipe, available map[string]int) int {
	count := 999

	for k, v := range recipe {
		availableCount, ok := available[k]
		if !ok {
			return 0
		}

		delta := availableCount / v
		if delta == 0 {
			return 0
		}

		if delta < count {
			count = delta
		}
	}

	return count
}

func main() {
	log.Printf("Start")

	res := Cakes(map[string]int{"flour": 500, "sugar": 200, "eggs": 1}, map[string]int{"flour": 1200, "sugar": 1200, "eggs": 5, "milk": 200})
	log.Printf("res: %v", res)
}
