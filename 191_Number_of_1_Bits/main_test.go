package _91_Number_of_1_Bits

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hammingWeight(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				num: 00000000000000000000000000001011,
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				num: 00000000000000000000000010000000,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := fmt.Sprintf("in %v want %d", tt.args.num, tt.want)
			assert.Equal(t, tt.want, hammingWeight(tt.args.num), m)
		})
	}
}
