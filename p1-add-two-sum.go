package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	var result [2]int
	var indexMap = make(map[int]int)
	for index, val := range nums {
		if i, ok := indexMap[val]; ok {
			if val*2 == target {
				result[0] = i
				result[1] = index
				return result[0:]
			}
			continue
		}
		indexMap[val] = index
	}
	for index, val := range nums {
		key := target - val
		if index2, ok := indexMap[key]; ok && key != val {
			result[0] = index
			result[1] = index2
			break
		}
	}
	return result[0:]
}

func main() {
	fmt.Println(twoSum([]int{3, 3}, 6))
}
