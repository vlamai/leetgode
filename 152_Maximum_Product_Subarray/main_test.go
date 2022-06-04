package _52_Maximum_Product_Subarray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{2, 3, -2, 4},
			},
			want: 6,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{-2, 0, -1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxProduct(tt.args.nums))
		})
	}
}
