package fifthres

func brackType(sym rune) int {
	switch sym {
	case '[':
		return 1
	case '{':
		return 1
	case '(':
		return 1
	case ']':
		return 2
	case '}':
		return 2
	case ')':
		return 2
	default:
		return 0
	}
}

var bracks = map[rune]rune{'[': ']', '{': '}', '(': ')'}

func isValid(s string) bool {
	var mas []rune
	for _, sym := range s {
		switch brackType(sym) {
		case 1:
			mas = append(mas, bracks[sym])
		case 2:
			if len(mas) > 0 && mas[len(mas)-1] == sym {
				mas = mas[:len(mas)-1]
			} else {
				return false
			}
		}
	}
	if len(mas) != 0 {
		return false
	}
	return true
}
