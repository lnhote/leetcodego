package main

import (
    "fmt"
    "sort"
)

func main() {
    // -4, -1, -1, 0, 1, 2
    fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // [ [-1, 0, 1], [-1, -1, 2] ]
    fmt.Println(threeSum([]int{-4,-2,1,-5,-4,-4,4,-2,0,4,0,-2,3,1,-5,0})) // [[-5,1,4],[-4,0,4],[-4,1,3],[-2,-2,4],[-2,1,1],[0,0,0]]
}
// https://leetcode.com/problems/3sum/
// The solution set must not contain duplicate triplets.
func threeSum(nums []int) [][]int {
    if len(nums) < 3 {
        return [][]int{}
    }
    if !sort.IntsAreSorted(nums) {
        sort.Ints(nums)
    }
    result := [][]int{}
    for i:=0; i<len(nums)-2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        a := nums[i]
        sumBC := 0 - a
        lo := i+1
        hi := len(nums)-1
        for lo < hi {
            b := nums[lo]
            c := nums[hi]
            if b + c < sumBC {
                lo++
            } else if b + c > sumBC {
                hi--
            } else {
                result = append(result, []int{a, b, c})
                for ;lo < hi && nums[lo] == nums[lo+1]; {
                    lo++
                }
                for ;lo < hi && nums[hi] == nums[hi-1]; {
                    hi--
                }
                lo++
                hi--
            }
        }
    }
    return result
}


