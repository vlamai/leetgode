package _41_Linked_List_Cycle

import (
	l "leetcode_patterns/linkedList"
)

func hasCycle(head *l.ListNode) bool {
	if head == nil {
		return false
	}
	slow := head.Next
	nf := head.Next
	if nf == nil {
		return false
	}
	fast := nf.Next
	for fast != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		nf := fast.Next
		if nf == nil {
			return false
		}
		fast = nf.Next
	}
	return false
}
