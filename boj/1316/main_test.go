package main

import (
	"fmt"
	"testing"
)

func Test_A(t *testing.T) {
	word := "word"

	for i := range word {
		fmt.Println(string(word[i]))
	}
}

func TestIsGroupWord(t *testing.T) {
	type args struct {
		w string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				w: "happy",
			},
			want: true,
		},
		{
			name: "",
			args: args{w: "new"},
			want: true,
		},
		{
			name: "",
			args: args{w: "year"},
			want: true,
		},
		{
			name: "",
			args: args{w: "abcabc"},
			want: false,
		},
		{
			name: "",
			args: args{w: "abab"},
			want: false,
		},
		{
			name: "",
			args: args{w: "a"},
			want: true,
		},
		{
			name: "",
			args: args{w: "aa"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsGroupWord(tt.args.w); got != tt.want {
				t.Errorf("IsGroupWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
