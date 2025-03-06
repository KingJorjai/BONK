package utils

import "strconv"

// Pluralize returns the plural suffix "s" if the given count is not equal to 1.
// If the count is 1, it returns an empty string.
//
// Parameters:
//   - count: an integer representing the count of items.
//
// Returns:
//   - A string that is either "s" for plural or an empty string for singular.
func Pluralize(count int) string {
	if count == 1 {
		return ""
	}
	return "s"
}

// CommaFormat takes an int64 number and returns a string representation of the number
// with commas inserted as thousand separators. For example, given the input
// 1234567, the function will return "1,234,567". Negative numbers are also
// handled correctly, with the minus sign preserved at the beginning of the string.
func CommaFormat(n int64) string {
	in := strconv.FormatInt(n, 10)
	numOfDigits := len(in)
	if n < 0 {
		numOfDigits-- // First character is the - sign (not a digit)
	}
	numOfCommas := (numOfDigits - 1) / 3

	out := make([]byte, len(in)+numOfCommas)
	if n < 0 {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ','
		}
	}
}
