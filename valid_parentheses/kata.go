package kata

// ValidParentheses https://www.codewars.com/kata/52774a314c2333f0a7000688/train/go
func ValidParentheses(parens string) bool {
	balance := 0

	for _, element := range []rune(parens) {
		if element == '(' {
			balance += 1
		} else if element == ')' {
			balance -= 1
		}

		if balance < 0 {
			return false
		}
	}

	return balance == 0
}
