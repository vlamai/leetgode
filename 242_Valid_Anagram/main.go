package _42_Valid_Anagram

func isAnagram(s string, t string) bool {
	return isAnagramWithoutMemoryAlloc(s, t)
}

func isAnagramWithoutMemoryAlloc(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	for _, sc := range s {
		for i, c := range t {
			if sc == c {
				t = t[:i] + t[i+1:]
				break
			}
		}
	}
	return t == ""
}

func isAnagramMap(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var set = make(map[uint8]int, 26)
	for i := range s {
		set[s[i]]++
		set[t[i]]--
	}
	for _, i := range set {
		if i < 0 {
			return false
		}
	}
	return true
}

func isAnagramArray(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	result := make([]int, 26)
	s0 := int('a')
	for i := 0; i < len(s); i++ {
		result[int(s[i])-s0]++
		result[int(t[i])-s0]--
	}
	for _, n := range result {
		if n != 0 {
			return false
		}
	}
	return true
}
