package _9_Combination_Sum

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type args struct {
	candidates []int
	target     int
}

type testCase struct {
	name string
	args args
	want [][]int
}

func tc(candidates []int, target int, want [][]int) testCase {
	name := fmt.Sprintf("candidates: %v, target: %d", candidates, target)
	return testCase{
		name: name,
		args: args{
			candidates: candidates,
			target:     target,
		},
		want: want,
	}
}
func Test_combinationSum(t *testing.T) {
	var tests []testCase

	tests = append(tests, tc([]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}))
	tests = append(tests, tc([]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}))
	tests = append(tests, tc([]int{2}, 1, [][]int{{}}))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, combinationSum(tt.args.candidates, tt.args.target))
		})
	}
}
