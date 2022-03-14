package main

import "fmt"

func main() {
	haystack := "abc"
	needle := "c"
	fmt.Print(strStr(haystack, needle))

}

func strStr(haystack string, needle string) int {
	if needle == "" && haystack == "" {
		return 0
	}
	if needle == haystack {
		return 0
	}
	for i := 0; i < len(haystack); i++ {
		if i+len(needle) <= len(haystack) {
			if haystack[i:len(needle)+i] == needle {
				return i
			}
		}
	}
	return -1
}
