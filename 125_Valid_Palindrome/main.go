package _25_Valid_Palindrome

func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	var noLet int
	for l < r {
		for IsLetter(rune(s[l])) == false && l < r {
			l++
			noLet++
		}
		for IsLetter(rune(s[r])) == false && l <= r {
			r--
			noLet++
		}
		lSym := low(s[l])
		rSym := low(s[r])
		if lSym != rSym {
			return false
		}
		l++
		r--
	}
	if noLet == len(s) {
		return false
	}
	return true
}

func IsLetter(r int32) bool {
	if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
		return false
	}
	return true
}

func low(s uint8) uint8 {
	if 'A' <= s && s <= 'Z' {
		s += 'a' - 'A'
	}
	return s
}
