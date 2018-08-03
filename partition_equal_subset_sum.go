package main

import "fmt"

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum % 2 == 1 {
		fmt.Println("fuck", sum)
		return false
	}
	sum = sum/2

	//dp[i][j] means whether the specific sum j can be gotten from the first i numbers.
	//	If we can pick such a series of numbers from 0-i whose sum is j, dp[i][j] is true, otherwise it is false.
	dp := make([][]bool, len(nums)+1)
	for i:=0; i<=len(nums); i++ {
		dp[i] = make([]bool, sum+1)
	}
	dp[0][0] = true
	for i, _ := range nums {
		dp[i][0] = true
	}
	for j := 1; j <= sum; j++ {
		dp[0][j] = false;
	}
	for i:= 1; i<len(nums); i++ {
		for j:= 1; j<=sum; j++ {
			dp[i][j] = dp[i-1][j];
			num := nums[i-1]
			if j >= num {
				fmt.Printf("dp[%d][%d]=%+v\n", i-1, j, dp[i-1][j])
				dp[i][j] = (dp[i])[j] || (dp[i-1])[j-num]
				fmt.Printf("sum(%d) >= num(%d): dp[%d][%d]=%+v\n", j, num, i, j, dp[i][j])
			}
		}
	}
	return dp[len(nums)-1][sum]
}

func main() {
	//fmt.Println(canPartition([]int{1,2,3,4,5,6})) // true
	//fmt.Println(canPartition([]int{2,2,2,3,3})) // true
	//fmt.Println(canPartition([]int{2,2,2,2,2,5,5})) // true
	//fmt.Println(canPartition([]int{6,4,4})) // false
	//fmt.Println(canPartition([]int{1,1,1,3,4})) // true
	//fmt.Println(canPartition([]int{1,5,11,5})) // true
	fmt.Println(canPartition([]int{1,2, 5})) // false
}