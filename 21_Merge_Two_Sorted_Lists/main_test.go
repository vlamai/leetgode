package _1_Merge_Two_Sorted_Lists

import (
	"github.com/stretchr/testify/assert"
	l "leetcode_patterns/linkedList"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 []int
		list2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				list1: []int{1, 2, 4},
				list2: []int{1, 3, 4},
			},
			want: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name: "Test 2",
			args: args{
				list1: []int{},
				list2: []int{},
			},
			want: []int{},
		},
		{
			name: "Test 3",
			args: args{
				list1: []int{1, 2, 4},
				list2: []int{},
			},
			want: []int{1, 2, 4},
		},
		{
			name: "Test 4",
			args: args{
				list1: []int{},
				list2: []int{1, 2, 4},
			},
			want: []int{1, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, mergeTwoLists(l.FromSlice(tt.args.list1), l.FromSlice(tt.args.list2)).ToSlice())
		})
	}
}
