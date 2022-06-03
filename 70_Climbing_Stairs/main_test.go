package _0_Climbing_Stairs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "Test 3",
			args: args{
				n: 4,
			},
			want: 5,
		},
		{
			name: "Test 4",
			args: args{
				n: 5,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, climbStairs(tt.args.n))
		})
	}
}
