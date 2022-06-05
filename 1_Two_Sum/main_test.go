package __Two_Sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_twoSum_clos(t *testing.T) {
	ts := func(nums []int, target int, want []int) func(t *testing.T) {
		return func(t *testing.T) {
			assert.Equal(t, want, twoSum(nums, target))
		}
	}
	t.Run("Test 1", ts([]int{2, 7, 11, 15}, 9, []int{0, 1}))
	t.Run("Test 2", ts([]int{2, 7, 11, 15}, 26, []int{2, 3}))
	t.Run("Test 3", ts([]int{2, 7, 11, 15}, -1, []int{}))

}

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "Test 2",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 26,
			},
			want: []int{2, 3},
		},
		{
			name: "Test 3",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: -1,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, twoSum(tt.args.nums, tt.args.target))
		})
	}
}
