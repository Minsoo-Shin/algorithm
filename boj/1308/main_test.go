package main

import (
	"testing"
)

func Test_IsLeapYear(t *testing.T) {
	tests := []struct {
		name string
		year int
		want bool
	}{
		{
			name: "400 배수",
			year: 2000,
			want: true,
		},
		{
			name: "400 배수",
			year: 1600,
			want: true,
		},
		{
			name: "100 배수",
			year: 2100,
			want: false,
		},
		{
			name: "100 배수",
			year: 2200,
			want: false,
		},
		{
			name: "4 배수",
			year: 2004,
			want: true,
		},
		{
			name: "4 배수",
			year: 2008,
			want: true,
		},
		{
			name: "4 배수 아님",
			year: 2009,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsLeapYear(tt.year)
			if got != tt.want {
				t.Fail()
			}
		})
	}
}
