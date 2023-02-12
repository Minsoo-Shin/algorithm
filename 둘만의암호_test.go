package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		s     string
		skip  string
		index int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s:     "aukks",
				skip:  "wbqd",
				index: 5,
			},
			want: "happy",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, solution(tt.args.s, tt.args.skip, tt.args.index), "solution(%v, %v, %v)", tt.args.s, tt.args.skip, tt.args.index)
		})
	}
}
