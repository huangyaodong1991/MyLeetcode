package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLenght := len(nums1) + len(nums2)
	if totalLenght & 0x1 != 0 {
		return float64(findKth(nums1,nums2,totalLenght/2 + 1))
	}else{
		return (float64(findKth(nums1,nums2,totalLenght/2)) + float64(findKth(nums1,nums2,totalLenght/2 + 1)))/2
	}
}

/**
两个排序数组发现第k个数的解法
利用性质：a + b = k, 则 nums1[a] 与nums2[b] 中较小的那个肯定比第k个数小
TODO:优化递归
 */
func findKth(nums1 []int, nums2 []int, k int) int  {
	l1 := len(nums1)
	l2 := len(nums2)
	if l1 > l2 {
		return findKth(nums2, nums1, k)
	}

	if l1 == 0 {
		return nums2[k-1]
	}

	if k == 1{
		if nums1[0] < nums2[0] { return nums1[0] } else {return nums2[0]}
	}

	var p1,p2 int
	if l1 < k/2 {
		p1 = l1
	}else{
		p1 = k/2
	}
	p2 = k - p1
	if nums1[p1-1] < nums2[p2-1] {
		return findKth(nums1[p1:],nums2,k-p1)
	}else if nums1[p1-1] == nums2[p2-1] {
		return nums1[p1-1]
	}else{
		return findKth(nums1,nums2[p2:],k-p2)
	}
}

func main() {
	nums1 := []int{1,8}
	nums2 := []int{4,5}
	result := findMedianSortedArrays(nums1, nums2)
	fmt.Println(result)
}