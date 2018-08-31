package main

import (
	"github.com/lnhote/leetcodego/actest"
)

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	sum = sum / 2
	n := len(nums)
	// dp[i][j] means whether the specific sum j can be gotten from the first i numbers.
	// i=0 means pick no number
	// if we can pick such a series of numbers from 0-i whose sum is j, dp[i][j] is true, otherwise it is false.
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, sum+1)
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = true
	}
	for i := 1; i <= n; i++ {
		num := nums[i-1]
		for j := 1; j <= sum; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= num {
				// fmt.Printf("dp[%d][%d]=%+v\n", i-1, j, dp[i-1][j])
				dp[i][j] = dp[i][j] || dp[i-1][j-num]
				// fmt.Printf("sum(%d) >= num(%d): dp[%d][%d]=%+v\n", j, num, i, j, dp[i][j])
			}
		}
	}
	return dp[n-1][sum]
}

func main() {
	actest.Equal(false, canPartition([]int{1, 2, 3, 4, 5, 6}))    // false
	actest.Equal(true, canPartition([]int{2, 2, 2, 3, 3}))       // true
	actest.Equal(true, canPartition([]int{2, 2, 2, 2, 2, 5, 5})) // true
	actest.Equal(false, canPartition([]int{6, 4, 4}))             // false
	actest.Equal(true, canPartition([]int{1, 1, 1, 3, 4}))       // true
	actest.Equal(true, canPartition([]int{1, 5, 11, 5}))         // true
	actest.Equal(false, canPartition([]int{1, 2, 5}))             // false
}
