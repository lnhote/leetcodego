package main

import (
    . "github.com/lnhote/leetcodego/node"
)

// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.

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
    var n1 = New([]int{2,4,3})
    var n2 = New([]int{5,6,4})
    Print(n1)
    Print(n2)
    Print(addTwoNumbers(n1, n2))
}