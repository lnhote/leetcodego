package main

import (
    "fmt"
)

// Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

func longestPalindrome(s string) string {
    if len(s) < 2 {
        return s
    }
    maxStr := ""
    for i:= 0; i<len(s); i++ {
        sub1 := expandAroundCenter(s, i, i)
        sub2 := expandAroundCenter(s, i, i+1)
        if len(sub1) > len(maxStr) {
            maxStr = sub1
        }
        if len(sub2) > len(maxStr) {
            maxStr = sub2
        }
    }
    return maxStr
}

func expandAroundCenter(s string, left, right int) string {
    if left < 0 || right > len(s) - 1 || s[left] != s[right] {
        return ""
    }
    L := left
    R := right
    for L >= 0 && R < len(s) && s[L] == s[R] {
        L--
        R++
    }
    // log.Printf("expandAroundCenter[%s][%d, %d]=[%d, %d]", s, left, right, L+1, R)
    return s[L+1:R]
}

func main() {
    fmt.Println(longestPalindrome("babad")) // bab
    fmt.Println(longestPalindrome("cbbd")) // bb
    fmt.Println(longestPalindrome("aaaabaaaaaaaa")) // aaaabaaaa
    fmt.Println(longestPalindrome("aaaabbaaaaaaaa")) // aaaabbaaaa
}