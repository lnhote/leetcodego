package main

import (
    "fmt"
)

// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
const (
    I =            1
    V =            5
    X =            10
    L =            50
    C =            100
    D =            500
    M =            1000
)

var (
    symboles = map[uint8]int{
        'I': I,
        'V': V,
        'X': X,
        'L': L,
        'C': C,
        'D': D,
        'M': M,
    }
)
//Given a roman numeral, convert it to an integer. Input is guaranteed to be within the range from 1 to 3999
func romanToInt(s string) int {
    if len(s) == 0 {
        return 0
    }
    result := 0
    var last = 0
    for i:= 0; i<len(s); i++ {
        if value, ok := symboles[s[i]]; ok {
            result += value
            if last != 0 && (value / last == 10 || value / last == 5) {
                result -= last * 2
            }
            last = value
        } else {
            return 0
        }

    }
    return result
}

func main() {
    fmt.Println(romanToInt("III")) // 3
    fmt.Println(romanToInt("IV")) // 4
    fmt.Println(romanToInt("IX")) // 9
    fmt.Println(romanToInt("LVIII")) // 58
    fmt.Println(romanToInt("MCMXCIV")) // 1994
}
