package _1_Merge_Two_Sorted_Lists

import l "leetcode_patterns/linkedList"

func mergeTwoLists(list1 *l.ListNode, list2 *l.ListNode) *l.ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var head, tail *l.ListNode
	if list1.Val < list2.Val {
		head, tail = list1, list1
		list1 = list1.Next
	} else {
		head, tail = list2, list2
		list2 = list2.Next
	}
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next, list1 = list1, list1.Next
			tail = tail.Next
		} else {
			tail.Next, list2 = list2, list2.Next
			tail = tail.Next
		}
	}
	if list1 != nil {
		tail.Next = list1
	}
	if list2 != nil {
		tail.Next = list2
	}
	return head
}
