package __Add_Two_Numbers

import (
	l "leetcode_patterns/linkedList"
)

func addTwoNumbers(l1 *l.ListNode, l2 *l.ListNode) *l.ListNode {
	extra := 0
	res := &l.ListNode{}
	head := res
	for {
		sum := extra
		if l1 != nil {
			sum += l1.Val
		}
		if l2 != nil {
			sum += l2.Val
		}

		extra = sum / 10

		res.Val = sum % 10

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		if l1 == nil && l2 == nil {
			if extra > 0 {
				res.Next = &l.ListNode{Val: extra}
			}
			break
		}
		res.Next = &l.ListNode{}
		res = res.Next

	}
	return head
}
