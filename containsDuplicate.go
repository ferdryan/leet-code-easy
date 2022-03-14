package main

import (
	"fmt"
)

func main() {
	// arr := []int{170, -368, 148, 672, 397, -629, -788, 192, 170, 309, -615, -237}

	arr := []int{0}
	fmt.Print(containsDuplicate(arr))
}
func containsDuplicate(nums []int) bool {
	amap := make(map[int]int)
	for _, v := range nums {
		if _, isExist := amap[v]; isExist {
			return true
		}
		amap[v] = v
	}
	fmt.Println(amap)
	return false
}

// func containsDuplicate(nums []int) bool {
// 	sort.Slice(nums, func(i, j int) bool {
// 		return nums[i] < nums[j]
// 	})
// 	fmt.Println(nums)
// 	for i := 0; i < len(nums)-1; i++ {
// 		if nums[i] == nums[i+1] {
// 			return true
// 		}
// 	}
// 	return false
// }

// func sort(arr []int, n int) []int {
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n-i-1; j++ {
// 			if arr[j] > arr[j+1] {
// 				arr[j], arr[j+1] = arr[j+1], arr[j]
// 			}
// 		}
// 	}
// 	return arr
// }
