package _41_Linked_List_Cycle

import (
	"github.com/stretchr/testify/assert"
	l "leetcode_patterns/linkedList"
	"testing"
)

func Test_hasCycle(t *testing.T) {
	withCicle := l.FromSliceWithCicle([]int{1, 2, 3, 4, 5}, 2)
	r := hasCycle(withCicle)
	assert.Equal(t, true, r)

	withoutCicle := l.FromSlice([]int{1, 2, 3, 4, 5})
	r = hasCycle(withoutCicle)
	assert.Equal(t, false, r)
}
