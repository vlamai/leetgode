package _03_Range_Sum_Query___Immutable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumArray_SumRange(t *testing.T) {
	n := Constructor([]int{-2, 0, 3, -5, 2, -1})

	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "Test 1",
			args: []int{0, 2},
			want: 1,
		},
		{
			name: "Test 2",
			args: []int{2, 5},
			want: -1,
		},
		{
			name: "Test 3",
			args: []int{0, 5},
			want: -3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, n.SumRange(tt.args[0], tt.args[1]))
		})
	}
}
