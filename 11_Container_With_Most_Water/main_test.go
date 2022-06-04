package _1_Container_With_Most_Water

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			want: 49,
		},
		{
			name: "Test 2",
			args: args{height: []int{1, 1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := fmt.Sprintf("in %v want %d", tt.args.height, tt.want)
			assert.Equal(t, tt.want, maxArea(tt.args.height), m)
		})
	}
}
