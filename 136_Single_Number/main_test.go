package _36_Single_Number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_singleNumber(t *testing.T) {
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
				nums: []int{2, 2, 1},
			},
			want: 1,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{4, 1, 2, 1, 2},
			},
			want: 4,
		},
		{
			name: "Test 3",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, singleNumber(tt.args.nums))
		})
	}
}
