package _5_3Sum

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{},
			},
			want: [][]int{},
		},
		{
			name: "Test 3",
			args: args{
				nums: []int{0},
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := fmt.Sprintf("in %v want out %v", tt.args.nums, tt.want)
			assert.Equal(t, tt.want, threeSum(tt.args.nums), m)
		})
	}
}
