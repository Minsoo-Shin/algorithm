package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				s: "leetcode",
			},
			want: 0,
		},
		{
			args: args{
				s: "loveleetcode",
			},
			want: 2,
		},
		{
			args: args{
				s: "aabb",
			},
			want: -1,
		},
		{
			args: args{
				s: "a",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := firstUniqChar(tt.args.s)

			assert.Equal(t, tt.want, got)

		})
	}
}

func Benchmark_firstUniqChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		firstUniqChar("leetcode")
	}
}
