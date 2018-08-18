package main

// https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/

// Given a sorted linked list, delete all duplicates such that each element appear only once.
import (
    . "github.com/lnhote/leetcodego/node"
    "github.com/lnhote/leetcodego/actest"
)

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    mark := head
    p := head.Next
    for p != nil {
        if p.Val == mark.Val {
            mark.Next = p.Next
        } else {
            mark = p
        }
        p = p.Next
    }
    return head
}

func main() {
    actest.Equal(New([]int{1}), deleteDuplicates(New([]int{1,1,1})))
    actest.Equal(New([]int{1,2}), deleteDuplicates(New([]int{1,1,2})))
    actest.Equal(New([]int{1,2,3}), deleteDuplicates(New([]int{1,1,2,3,3})))
}