package main

import "fmt"

func GetMiddle(s string) string {
	address := len(s) / 2
	end := address + 1
	if len(s)%2 == 0 {
		address--
	}
	return s[address:end]
}

func main() {
	//res := GetMiddle("testing")
	res := GetMiddle("test")
	fmt.Printf("Res: %s\n", res)
}
