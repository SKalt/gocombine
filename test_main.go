package gocombine

import (
	"testing"
	// "github.com/stretchr/testify"
)

func TestCombine(t *testing.T) {
	combos := Combinations(
		[]string{"a", "b"}, []string{"x", "y"}, []string{"1", "2", "3"})
	expected := [][]string{
		{"a", "x", "1"},
		{"a", "x", "2"},
		{"a", "x", "3"},
		{"a", "y", "1"},
		{"a", "y", "2"},
		{"a", "y", "3"},
		{"b", "x", "1"},
		{"b", "x", "2"},
		{"b", "x", "3"},
		{"b", "y", "1"},
		{"b", "y", "2"},
		{"b", "y", "3"},
	}
	for i, combo := range combos {
		for j, str := range combo {

		}
	}
}
