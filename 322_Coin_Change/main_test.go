package _22_Coin_Change

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				coins:  []int{1, 2, 5},
				amount: 11,
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				coins:  []int{2},
				amount: 3,
			},
			want: -1,
		},
		{
			name: "Test 3",
			args: args{
				coins:  []int{1},
				amount: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := fmt.Sprintf("in %v | %d want %d", tt.args.coins, tt.args.amount, tt.want)
			assert.Equal(t, tt.want, coinChange(tt.args.coins, tt.args.amount), m)
		})
	}
}
