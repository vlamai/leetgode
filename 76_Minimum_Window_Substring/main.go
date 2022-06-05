package _6_Minimum_Window_Substring

// Given two strings s and t of lengths m and n respectively,
// return the minimum window substring of s such that every character in t (including duplicates)
// is included in the window. If there is no such substring, return the empty string "".
//
// The testcases will be generated such that the answer is unique.
//
// A substring is a contiguous sequence of characters within the string.

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	numOfChars := make(map[uint8]int)
	for _, i2 := range t {
		char := uint8(i2)
		if _, ok := numOfChars[char]; ok {
			numOfChars[char]++
		} else {
			numOfChars[char] = 1
		}
	}

	for _, i := range s {
		char := uint8(i)
		if _, ok := numOfChars[char]; ok {
			numOfChars[char]--
		}
	}
	for _, i := range numOfChars {
		if i > 0 {
			return ""
		}
	}
	if len(s) == len(t) {
		return s
	}
	l := 0

	for {
		char := s[l]
		if _, ok := numOfChars[char]; ok {
			numOfChars[char]++
			if numOfChars[char] > 0 {
				break
			} else {
				l++
			}
		} else {
			l++
		}

	}
	r := len(s) - 1
	for {
		char := s[r]
		if _, ok := numOfChars[char]; ok {
			numOfChars[char]++
			if numOfChars[char] > 0 {
				break
			} else {
				r--
			}
		} else {
			r--
		}
	}
	return s[l : r+1]
}
