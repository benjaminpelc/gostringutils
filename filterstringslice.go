package gostringutils

// FilterStringSlice
// Filter the string slice ss using predicate fn.
//
// Move this to gostringutils package.
func FilterStringSlice(ss []string, fn func(string) bool) []string {
	var selected []string
	for _, s := range ss {
		if fn(s) {
			selected = append(selected, s)
		}
	}
	return selected
}
