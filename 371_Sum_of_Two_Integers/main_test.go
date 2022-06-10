package _71_Sum_of_Two_Integers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				a: 1,
				b: 2,
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				a: 3,
				b: 2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getSum(tt.args.a, tt.args.b))
		})
	}
}
