package kata

// ValidParentheses https://www.codewars.com/kata/52774a314c2333f0a7000688/train/go
func ValidParentheses(parens string) bool {
	balance := 0

	for _, element := range parens {
		if element == '(' {
			balance++
		} else if element == ')' {
			balance--
		}

		if balance < 0 {
			return false
		}
	}

	return balance == 0
}
