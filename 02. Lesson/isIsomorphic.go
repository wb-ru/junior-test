package isIsomorphic

import (
	"strings"
)

func isIsomorphic(s string, t string) bool {
	m := make(map[string]string)
	lettersFirst := strings.Split(s, "")
	lettersSecond := strings.Split(t, "")

	for i := 0; i < len(lettersFirst); i++ {
		if m[lettersFirst[i]] == "" {
			m[lettersFirst[i]] = lettersSecond[i]
		} else if m[lettersFirst[i]] != lettersSecond[i] {
			return false
		}
	}

	return true
}
