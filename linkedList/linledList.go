package linkedList

type ListNode struct {
	val  int
	next *ListNode
}

func FromSliceWithCicle(nums []int, pos int) *ListNode {
	head := &ListNode{
		val: nums[0],
	}
	tmp := head
	var cicleNode *ListNode
	for i, num := range nums[1:] {
		if i == pos {
			cicleNode = tmp
		}
		n := &ListNode{val: num}
		tmp.next = n
		tmp = n
	}
	tmp.next = cicleNode
	return head
}

func FromSlice(nums []int) *ListNode {
	head := &ListNode{
		val: nums[0],
	}
	tmp := head
	for _, num := range nums[1:] {
		n := &ListNode{val: num}
		tmp.next = n
		tmp = n
	}
	return head
}
