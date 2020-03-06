package secondres

func isIsomorphic(s string, t string) bool {
	isomorph := make(map[string]string)
	for i := 0; i < len(s); i++ {
		first := string(s[i])
		second := string(t[i])
		_, ok := isomorph[first]
		if ok == false {
			isomorph[first] = second
		} else if isomorph[first] != second {
			return false
		}
	}
	return true
}
