// Package main provides string utilities
package gostringutils

// StringInSlice
// Check if the slice of strings, list, contains the string
// str. Return true if str is present and false otherwise
func StringInSlice(str string, list []string) bool {
	for _, val := range list {
		if str == val {
			return true
		}
	}
	return false
}