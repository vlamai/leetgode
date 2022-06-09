package _42_Valid_Anagram

func isAnagram(s string, t string) bool {
	return isAnagramMap(s, t)
}

func isAnagramMap(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var set = make(map[int32]int, 26)
	for _, c := range s {
		if _, ok := set[c]; ok {
			set[c]++
		} else {
			set[c] = 1
		}
	}
	for _, c := range t {
		if _, ok := set[c]; ok {
			set[c]--
			if set[c] < 0 {
				return false
			}
		} else {
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
