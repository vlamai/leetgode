package _6_Minimum_Window_Substring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{
				s: "ADOBECODEBANC",
				t: "ABC",
			},
			want: "BANC",
		},
		{
			name: "Test 2",
			args: args{
				s: "a",
				t: "a",
			},
			want: "a",
		},
		{
			name: "Test 3",
			args: args{
				s: "a",
				t: "aa",
			},
			want: "",
		},
		{
			name: "Test 4",
			args: args{
				s: "cabwefgewcwaefgcf",
				t: "cae",
			},
			want: "cwae",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, minWindow(tt.args.s, tt.args.t))
		})
	}
}
