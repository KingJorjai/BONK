package utils_test

import (
	"testing"

	"github.com/KingJorjai/BONK/pkg/utils"
)

func TestPluralize(t *testing.T) {
	tests := []struct {
		count    int
		expected string
	}{
		{1, ""},
		{0, "s"},
		{2, "s"},
		{100, "s"},
	}

	for _, test := range tests {
		result := utils.Pluralize(test.count)
		if result != test.expected {
			t.Errorf("Pluralize(%d) = %s; want %s", test.count, result, test.expected)
		}
	}
}

func TestCommaFormat(t *testing.T) {
	tests := []struct {
		number   int64
		expected string
	}{
		{4294967295, "4,294,967,295"},
		{1234567, "1,234,567"},
		{-1234567, "-1,234,567"},
		{1000, "1,000"},
		{0, "0"},
		{-1000, "-1,000"},
		{-4294967295, "-4,294,967,295"},
	}

	for _, test := range tests {
		result := utils.CommaFormat(test.number)
		if result != test.expected {
			t.Errorf("CommaFormat(%d) = %s; want %s", test.number, result, test.expected)
		}
	}
}
