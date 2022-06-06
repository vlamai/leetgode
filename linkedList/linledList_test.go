package linkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedListFromSlice(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	l := FromSlice(ints)
	var i int
	for l != nil {
		assert.Equal(t, ints[i], l.Val)
		i++
		l = l.Next
	}
	assert.Equal(t, len(ints), i)
}

func TestListNode_FromSliceWithCicle(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	l := FromSliceWithCicle(ints, 2)
	for i := 0; i < len(ints); i++ {
		l = l.Next
	}
	assert.Equal(t, l.Val, ints[2])
}

func TestListNode_FromSliceWithCicle2(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	l := FromSliceWithCicle(ints, 1)
	for i := 0; i < len(ints); i++ {
		l = l.Next
	}
	assert.Equal(t, l.Val, ints[1])
}

func TestToSlice(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	l := FromSlice(ints)
	assert.Equal(t, ints, l.ToSlice())
}

func TestListNode_String(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	l := FromSlice(ints)
	assert.Equal(t, "1 -> 2 -> 3 -> 4 -> 5 -> nil", l.String())
}
