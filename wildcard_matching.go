package main

import (
    "github.com/lnhote/leetcodego/actest"
)

// https://leetcode.com/problems/wildcard-matching/description/

// '?' Matches any single character.
// '*' Matches any sequence of characters (including the empty sequence).
func isMatch(s string, p string) bool {
    dp := make([][]bool, len(s)+1)
    for i := 0; i <= len(s); i++ {
        dp[i] = make([]bool, len(p)+1)
    }
    dp[0][0] = true
    for i := 1; i <= len(s); i++ {
        dp[i][0] = false
    }
    for j := 1; j <= len(p); j++ {
        dp[0][j] = p[j-1] == '*' && dp[0][j-1]
    }
    for i := 1; i <= len(s); i++ {
        for j := 1; j <= len(p); j++ {
            if p[j-1] == '*' {
                // match 1, match more than 1, or match none
                dp[i][j] = dp[i-1][j-1] || dp[i-1][j] || dp[i][j-1]
            } else if p[j-1] == '?' {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = dp[i-1][j-1] && s[i-1] == p[j-1]
            }
            // fmt.Printf("dp[%d][%d]=%v\n", i, j, dp[i][j])
        }
    }
    return dp[len(s)][len(p)]
}

func main() {
    actest.False(isMatch("", "?"))
    actest.False(isMatch("aab", "c*a*b"))
    actest.False(isMatch("aa", "a"))
    actest.True(isMatch("aa", "a*"))
    actest.True(isMatch("aa", "*"))
    actest.True(isMatch("a", "a*"))
    actest.True(isMatch("aaaaa", "a*"))
    actest.False(isMatch("cb", "?a"))
    actest.True(isMatch("adceb", "*a*b"))
    actest.False(isMatch("acdcb", "a*c?b"))
    actest.True(isMatch("abbabbbaabaaabbbbbabbabbabbbabbaaabbbababbabaaabbab", "*aabb***aa**a******aa*"))
}
