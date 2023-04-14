package main

import (
	"testing"
)

func TestHashFunctionShouldReturnHashForTheGivenInput(t *testing.T) {
	type test struct {
		name   string
		input  string
		output int
	}

	tests := []test{
		{
			name:   "should return hash value 100",
			input:  "input",
			output: 100,
		},
		{
			name:   "should return hash value 332",
			input:  "abc",
			output: 332,
		},
		{
			name:   "should return hash value 249",
			input:  "xyz",
			output: 249,
		},
	}

	testHash := NewHash()
	for _, tt := range tests {
		got := testHash.hash(tt.input, 1000, 1)

		if got != tt.output {
			t.Errorf("got: %d, wanted: %d", got, tt.output)
		}
	}

}
