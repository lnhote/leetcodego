package node

import (
    "fmt"
    "strings"
    "strconv"
)

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
    Val int
    Next *ListNode
}

func (l *ListNode) String() string {
    arr := []string{}
    for l != nil {
        arr = append(arr, strconv.Itoa(l.Val))
        l = l.Next
    }
    return strings.Join(arr, "->")
}

func New(arr []int) *ListNode {
    l := &ListNode{}
    p := l
    for i, val := range arr {
        p.Val = val
        if i != len(arr)-1 {
            p.Next = &ListNode{}
        }
        p = p.Next
    }
    return l
}

func Print(head *ListNode) {
    for head != nil {
        fmt.Print(head.Val)
        head = head.Next
    }
    fmt.Print("\n")
}