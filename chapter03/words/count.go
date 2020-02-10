package words

import (
	"strings"
)

// CountWords counts the number of words in the specified
// String and returns the count.
func CountWords(text string) int {
	return len(strings.Fields(text))
}
