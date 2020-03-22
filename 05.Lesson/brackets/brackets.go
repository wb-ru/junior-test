package brackets

func popStack(s []int32) ([]int32, int32) {
	return s[:len(s)-1], s[len(s)-1]
}

func isValid(s string) bool {
	var stack []int32
	var pop int32
	for _, item := range s {
		switch item {
		case '(':
			stack = append(stack, item)
		case '{':
			stack = append(stack, item)
		case '[':
			stack = append(stack, item)
		default:
			if len(stack) == 0 {
				return false
			}
			stack, pop = popStack(stack)
			switch pop {
			case '(':
				if item != ')' {
					return false
				}
			case '{':
				if item != '}' {
					return false
				}
			case '[':
				if item != ']' {
					return false
				}

			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
