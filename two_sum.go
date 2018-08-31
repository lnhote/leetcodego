package main

import (
    "github.com/lnhote/leetcodego/actest"
)

func twoSum(nums []int, target int) []int {
    dic := make(map[int]int, len(nums))
    var r1, r2 int
    for _, v := range nums {
        counterpart := target - v
        if exist, ok := dic[counterpart]; exist == 1 && ok {
            return []int{v, counterpart}
        } else {
            dic[v] = 1
        }
    }
    return []int{r1, r2}
}

func main() {
    actest.EqualValues([]int{4,2}, twoSum([]int{1,2,3,4,5}, 6))
}