package daily

// This problem was asked by Facebook.

// Given a string of round, curly, and square open and closing brackets, return whether the brackets are balanced (well-formed).
// For example, given the string "([])[]({})", you should return true.
// Given the string "([)]" or "((()", you should return false.

func pop(in []rune) rune {
	if len(in) == 0 {
		return ' '
	}

	out := in[len(in)-1]
	in = in[:len(in)-1]
	return out
}

func isBalanced(in string) bool {
	var stack []rune
	for _, b := range in {
		switch b {
		case '(', '[', '{':
			stack = append(stack, b)
			continue
		}

		cur := pop(stack)
		if b == ')' && cur == '(' ||
			b == ']' && cur == '[' ||
			b == '}' && cur == '{' {
			continue
		}

		return false
	}

	if len(stack) == 0 {
		return true
	}

	return false
}
