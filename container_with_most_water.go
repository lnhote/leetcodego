package main

import "fmt"

// https://leetcode.com/problems/container-with-most-water/description/

func maxArea(height []int) int {
	max:=0
	lo:=0
	hi:=len(height)-1
	for lo < hi {
		min := 0
		if height[lo]>=height[hi] {
			min = height[hi]
		} else {
			min = height[lo]
		}
		if min*(hi-lo) >= max {
			max = min*(hi-lo)
		}
		if height[lo]>=height[hi] {
			hi--
		} else {
			lo++
		}
	}
	return max
}

func main() {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7})) // 49
}