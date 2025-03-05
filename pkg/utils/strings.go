package utils

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
