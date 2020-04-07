package brackets

var (
	stack []int32
)

func isOpenBracket(value int32) bool {
	switch value {
	case '(', '[', '{':
		return true
	}
	return false
}

func isMatchBracket(closeBracket int32) bool {
	openBracket := stack[len(stack)-1]
	if openBracket == '(' {
		if closeBracket != ')' {
			return false
		}
	} else if openBracket == '[' {
		if closeBracket != ']' {
			return false
		}
	} else if openBracket == '{' {
		if closeBracket != '}' {
			return false
		}
	}
	return true
}

func isValid(sequence string) bool {
	if len(sequence) == 0 {
		return true
	}
	stack = make([]int32, 0)
	for _, value := range sequence {
		if isOpenBracket(value) {
			stack = append(stack, value)
		} else {
			if len(stack) == 0 {
				return false
			}
			if isMatchBracket(value) {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
