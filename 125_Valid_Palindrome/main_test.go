package _25_Valid_Palindrome

import "testing"

func Test_isPalindrome(t *testing.T) {
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
				s: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				s: "race a car",
			},
			want: false,
		},
		{
			name: "Test 3",
			args: args{
				s: " ",
			},
			want: true,
		},
		{
			name: "Test 4",
			args: args{
				s: ",,",
			},
			want: true,
		},
		{
			name: "Test 5",
			args: args{
				s: "0P",
			},
			want: true,
		},
		{
			name: "Test 6",
			args: args{
				s: ".,",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
