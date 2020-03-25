package isomorphic

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
					return false}
	s1:=[]rune(s)
	t1:=[]rune(t)

	m:= make (map[rune]rune)
	for i:=0; i<len(s); i++ {
		_,flag := m[s1[i]]
		if !flag { m[s1[i]]=t1[i]
		} else {
			if m[s1[i]]!=t1[i] {
				return false
			}
		}
	}
	return true
}
