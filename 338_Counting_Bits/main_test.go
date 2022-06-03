package _38_Counting_Bits

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countBits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				n: 2,
			},
			want: []int{0, 1, 1},
		},
		{
			name: "Test 2",
			args: args{
				n: 5,
			},
			want: []int{0, 1, 1, 2, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, countBits(tt.args.n))
		})
	}
}
