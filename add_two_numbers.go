package main

import (
    "fmt"
)

// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
    var prev *ListNode = nil
    var head *ListNode = nil
    var p *ListNode
    for (l1 != nil || l2 != nil) {
        sum := 0
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }
        sum += carry
        if sum >= 10 {
            carry = 1
            sum = sum - 10
        } else {
            carry = 0
        }
        p = &ListNode{sum, nil}
        if prev != nil {
            prev.Next = p
        } else {
            head = p
        }
        prev = p
    }
    if carry > 0 && prev != nil {
        prev.Next = &ListNode{carry, nil}
    }
    return head
}

func main() {
    var n1 = &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
    var n2 = &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
    printLNumber(n1)
    printLNumber(n2)
    printLNumber(addTwoNumbers(n1, n2))
}

func printLNumber(l *ListNode) {
    for l != nil {
        fmt.Print(l.Val)
        l = l.Next
    }
    fmt.Print("\n")
}