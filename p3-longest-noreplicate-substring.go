package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var mapping [257]int
	var start,current,maxLength int
	if len(s) == 0 {
		return 0
	}
	for key, value := range s {
		index := int(value)
		if i :=mapping[index];i != 0 {
			lastStart := start
			start = i+1
			current = key - start + 1
			if current > maxLength {
				maxLength = current
			}
			for ;lastStart < start;lastStart++ {
				tmp := int(s[lastStart])
				mapping[tmp] = 0
			}
		}
		mapping[index] = key
	}
	return maxLength
}

func main(){
	longestSubstring := lengthOfLongestSubstring("bbbbbbb")
	fmt.Println(longestSubstring)
}
