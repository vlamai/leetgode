package _42_Valid_Anagram

func isAnagram(s string, t string) bool {
	return isAnagramMap(s, t)
}

func isAnagramMap(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var set = make(map[int32]int)
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
