package main

import "testing"

func TestValidElem(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{s: "(())("},
			want: 2,
		},
		{
			name: "",
			args: args{s: "())"},
			want: 1,
		},
		{
			name: "",
			args: args{s: ")((("},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidElem(tt.args.s); got != tt.want {
				t.Errorf("ValidElem() = %v, want %v", got, tt.want)
			}
		})
	}
}
