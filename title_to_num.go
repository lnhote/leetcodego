package main

import "github.com/lnhote/leetcodego/actest"

// https://leetcode.com/problems/excel-sheet-column-number/description/

func titleToNumber(s string) int {
	if len(s) == 0 {
		return 0
	}
	sum := 0
	for i := 0; i < len(s); i++ {
		num := int(s[i] - 'A' + 1)
		sum = sum*26 + num
	}
	return sum
}

func main() {
	actest.Equal(titleToNumber("A"), 1)
	actest.Equal(titleToNumber("B"), 2)
	actest.Equal(titleToNumber("Z"), 26)
	actest.Equal(titleToNumber("AA"), 27)
	actest.Equal(titleToNumber("BA"), 53)
	actest.Equal(titleToNumber("ZY"), 701)
}
