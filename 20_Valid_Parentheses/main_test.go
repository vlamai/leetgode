package _0_Valid_Parentheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			name: "Test 1",
			args: args{
				s: "(]",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isValid(tt.args.s))
		})
	}
}
