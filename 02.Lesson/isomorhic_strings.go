package isomorphic_strings

func isIsomorphic(s string, t string) bool {
	ls, lt := len(s), len(t)
	if ls != lt {
		return false // strings are of differents lengths
	}

	var visited [2048]int
	m := make(map[int32]uint8)
	for i, item := range s {
		_, ok := m[item]
		if !ok {
			if visited[t[i]] != 0 {
				return false
			}

			visited[t[i]] = 1
			m[item] = t[i]
		} else if m[item] != t[i] {
			return false
		}

	}

	return true
}
