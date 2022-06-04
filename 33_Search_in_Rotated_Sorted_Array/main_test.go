package _3_Search_in_Rotated_Sorted_Array

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			want: 4,
		},
		{
			name: "Test 2",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			want: -1,
		},
		{
			name: "Test 3",
			args: args{
				nums:   []int{1},
				target: 0,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := fmt.Sprintf("in %v find %d must be at %d\n", tt.args.nums, tt.args.target, tt.want)
			assert.Equal(t, tt.want, search(tt.args.nums, tt.args.target), m)
		})
	}
}
