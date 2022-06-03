package linkedList

type ListNode struct {
	Val  int
	Next *ListNode
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
