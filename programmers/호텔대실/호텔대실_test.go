package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_호텔(t *testing.T) {
	assert.Equal(t, 1, 둘만의암호([][]string{
		{
			"09:10",
			"10:10",
		},
		{
			"10:20",
			"18:00",
		},
	}))

}
