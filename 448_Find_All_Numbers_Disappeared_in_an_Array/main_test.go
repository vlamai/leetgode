package _48_Find_All_Numbers_Disappeared_in_an_Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findDisappearedNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
			},
			want: []int{5, 6},
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{1, 1},
			},
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findDisappearedNumbers(tt.args.nums))
		})
	}
}
