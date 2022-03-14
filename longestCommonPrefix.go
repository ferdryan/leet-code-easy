package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	var str string
	for _, stemp := range strs {
		str = strs[0]
		for i := 1; i <= len(stemp) && i < len(str)-1; i++ {
			if str[0:i] != stemp[0:i] {
				str = stemp[0 : i-1]
			}
		}
	}
	return str
}

func main() {
	strs := []string{"ab", "a"}
	fmt.Println(longestCommonPrefix(strs))
}
