// Package main provides string utilities
package gostringutils

import "regexp"

var trailNonWordChars = regexp.MustCompile(`\W+$`)
var leadingNowWordExceptHashOrAt = regexp.MustCompile(`^[\:\+\-\)\(\"]+`)

// StripTrailingNonWordChars
// Take string str and strip all trailing non-word characters
func StripTrailingNonWordChars(str string) string {
	return trailNonWordChars.ReplaceAllString(str, "")
}

// StripLeadingNonWordsLeaveAtAndHash
func StripLeadingNonWordsLeaveAtAndHash(str string) string {
	return leadingNowWordExceptHashOrAt.ReplaceAllString(str, "")
}

// TweetWordStrip
// Remove all trailing and leading non-word characters unless
// the leading charachter is a # or an @
func TweetWordStrip(str string) string {
	str = StripTrailingNonWordChars(str)
	return StripLeadingNonWordsLeaveAtAndHash(str)
}
