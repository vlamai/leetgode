package linkedList

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// ToSlice Convert linked list to a slice
func (h *ListNode) ToSlice() []int {
	var nums []int
	head := h
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return nums
}

// Print print linked list
func (h *ListNode) String() string {
	if h == nil {
		return "nil"
	}
	var s = ""
	head := h
	for head != nil {
		s += fmt.Sprintf("%d -> ", head.Val)
		head = head.Next
	}
	s += "nil"
	return s
}

func FromSliceWithCicle(nums []int, pos int) *ListNode {
	head := &ListNode{
		Val: nums[0],
	}
	tmp := head
	var cicleNode *ListNode
	for i, num := range nums[1:] {
		if i == pos {
			cicleNode = tmp
		}
		n := &ListNode{Val: num}
		tmp.Next = n
		tmp = n
	}
	tmp.Next = cicleNode
	return head
}

func FromSlice(nums []int) *ListNode {
	head := &ListNode{
		Val: nums[0],
	}
	tmp := head
	for _, num := range nums[1:] {
		n := &ListNode{Val: num}
		tmp.Next = n
		tmp = n
	}
	return head
}
