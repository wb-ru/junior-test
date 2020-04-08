package isomorphic

func isIsomorphic(s string, t string) bool {
	match := make(map[uint8]uint8)
	used := make(map[uint8]bool)
	for i := range s {
		runeS := s[i]
		runeT := t[i]
		if value, ok := match[runeS]; ok {
			if value != runeT {
				return false
			}
		} else {
			if used[runeT] {
				return false
			}
			match[runeS] = runeT
			used[runeT] = true
		}
	}
	return true
}
