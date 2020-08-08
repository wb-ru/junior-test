package lesson2

func isIsomorphic(s string, t string) bool {

	m := make(map[int]int)

	for i, v := range s {
		keyMap, ok := m[int(v)]
		if !ok {
			m[int(v)] = int(t[i])
		} else {
			if int(t[i]) != keyMap {
				return false
			}
		}
	}
	return true
}
