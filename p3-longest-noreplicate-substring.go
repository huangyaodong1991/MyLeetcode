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
			start = i
			current = key - start + 1
			for ;lastStart < start;lastStart++ {
				tmp := int(s[lastStart])
				mapping[tmp] = 0
			}
		}else{
			current += 1
			if current > maxLength {
				maxLength = current
			}
		}
		//此处不可使mapping中对应的位置为0，否则会与第一个字符的位置混淆，如aaaa的情况
		mapping[index] = key+1
	}
	return maxLength
}

func main(){
	longestSubstring := lengthOfLongestSubstring("abcded")
	fmt.Println(longestSubstring)
}
