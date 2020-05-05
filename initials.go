package initials

import "strings"

// Find gets the first letter of each word, returning the 'intiials'.
func Find(s string) (i string) {
	s = strings.ReplaceAll(s, "  ", " ")
	x := strings.Split(s, " ")

	for _, y := range x {
		i += strings.Title(strings.ToLower(string(y[0])))
	}

	return i
}
