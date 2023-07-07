package main

import "fmt"

func sort(source map[string][]string, lib string, set map[string]struct{}) {
	deps, ok := source[lib]
	if !ok {
		if _, ok := set[lib]; !ok {
			fmt.Printf("%s ", lib)
			set[lib] = struct{}{}
		}
		return
	}

	for _, dep := range deps {
		sort(source, dep, set)
	}
	if _, ok := set[lib]; !ok {
		fmt.Printf("%s ", lib)
		set[lib] = struct{}{}
	}
}

func Entrypoint(source map[string][]string) {
	set := make(map[string]struct{})

	for k, _ := range source {
		//fmt.Printf("ENTRY: %s; set: %v\n", k, set)
		sort(source, k, set)
	}
}

func swap(list []int) []int {
	pos := 0

	for i, el := range list {
		if el == 0 {
			for j := i + pos; j < len(list); j++ {
				if list[j] != 0 {
					list[i], list[j] = list[j], list[i]
					pos = j - i
					break
				}
			}
		}
	}

	return list
}

func main() {

	deps := map[string][]string{
		"tensorflow": {"nvcc", "gpu", "linux"},
		"nvcc":       {"linux"},
		"linux":      {"core"},
		"mylib":      {"tensorflow"},
		"mylib2":     {"requests"},
	}

	fmt.Printf("Start\n")

	Entrypoint(deps)
	fmt.Println()

	//resSwap := swap([]int{7, 3, 0, 0, 0, 0, 0, 0, 0, 2, 4, 0, 5, 19})
	resSwap := swap([]int{7, 3, 2, 4, 5, 0, 19})
	fmt.Printf("resSwap: %v\n", resSwap)
	// [7 3 2 4 5 19 0 0 0 0 0 0 0 0]

}
