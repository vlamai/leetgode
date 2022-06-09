package _42_Valid_Anagram

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				s: "rat",
				t: "car",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isAnagram(tt.args.s, tt.args.t))
		})
	}
}

func StringWithCharset(length int, charset string) string {
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func benchmark_isAnagramMap(l int, b *testing.B) {
	s, t := GetStrings(l)
	for i := 0; i < b.N; i++ {
		isAnagramMap(s, t)
	}
}
func Benchmark_isAnagramMap10(b *testing.B)    { benchmark_isAnagramMap(10, b) }
func Benchmark_isAnagramMap100(b *testing.B)   { benchmark_isAnagramMap(100, b) }
func Benchmark_isAnagramMap1000(b *testing.B)  { benchmark_isAnagramMap(1000, b) }
func Benchmark_isAnagramMap10000(b *testing.B) { benchmark_isAnagramMap(10000, b) }

func benchmark_isAnagramArray(l int, b *testing.B) {
	s, t := GetStrings(l)
	for i := 0; i < b.N; i++ {
		isAnagramArray(s, t)
	}
}

func Benchmark_isAnagramArray10(b *testing.B)    { benchmark_isAnagramArray(10, b) }
func Benchmark_isAnagramArray100(b *testing.B)   { benchmark_isAnagramArray(100, b) }
func Benchmark_isAnagramArray1000(b *testing.B)  { benchmark_isAnagramArray(1000, b) }
func Benchmark_isAnagramArray10000(b *testing.B) { benchmark_isAnagramArray(10000, b) }

func GetStrings(l int) (string, string) {
	const charset = "abcdefghijklmnopqrstuvwxyz"
	s := StringWithCharset(l, charset)

	inRune := []rune(s)
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	t := string(inRune)
	return s, t
}
