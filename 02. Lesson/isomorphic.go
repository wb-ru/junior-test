package isomorphic

func isIsomorphic(s string, t string) bool {

	strMap := make(map[int]int)

	for i, value := range s {
		mapValue, inMap := strMap[int(value)]
		if !inMap {
			strMap[int(value)] = int(t[i])
		} else {

			if int(t[i]) != mapValue {
				return false
			}
		}

	}

	return true
}
