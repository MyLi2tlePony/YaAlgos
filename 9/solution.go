package main

func isValid(str string) bool {
	var stack []byte

	for i := range str {
		if len(stack) > 0 && isPair(str[i], stack[len(stack)-1]) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, str[i])
		}
	}

	return len(stack) == 0
}

func isPair(closed, opened byte) bool {
	switch closed {
	case ')':
		return opened == '('
	case '}':
		return opened == '{'
	case ']':
		return opened == '['
	}

	return false
}
