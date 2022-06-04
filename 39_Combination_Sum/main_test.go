package _9_Combination_Sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test 1",
			args: args{
				candidates: []int{2, 3, 6, 7},
				target:     7,
			},
			want: [][]int{
				{2, 2, 3},
			},
		},
		{
			name: "Test 2",
			args: args{
				candidates: []int{2, 3, 5},
				target:     8,
			},
			want: [][]int{
				{2, 2, 2, 2},
				{2, 3, 3},
				{3, 5},
			},
		},
		{
			name: "Test 3",
			args: args{
				candidates: []int{2},
				target:     1,
			},
			want: [][]int{{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, combinationSum(tt.args.candidates, tt.args.target))
		})
	}
}