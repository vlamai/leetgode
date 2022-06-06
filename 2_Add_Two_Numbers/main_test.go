package __Add_Two_Numbers

import (
	"github.com/stretchr/testify/assert"
	l "leetcode_patterns/linkedList"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				l1: []int{2, 4, 3},
				l2: []int{5, 6, 4},
			},
			want: []int{7, 0, 8},
		},
		{
			name: "Test 2",
			args: args{
				l1: []int{5},
				l2: []int{5},
			},
			want: []int{0, 1},
		},
		{
			name: "Test 3",
			args: args{
				l1: []int{9, 9, 9, 9, 9, 9, 9},
				l2: []int{9, 9, 9, 9},
			},
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
		{
			name: "Test 4",
			args: args{
				l1: []int{0},
				l2: []int{0},
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, addTwoNumbers(l.FromSlice(tt.args.l1), l.FromSlice(tt.args.l2)).ToSlice())
		})
	}
}
