package _44_Backspace_String_Compare

func backspaceCompare(s string, t string) bool {
	return getString(s) == getString(t)
}

func getString(s string) (r string) {
	var d int
	for i := len(s) - 1; i >= 0; i-- {
		l := s[i]
		if l == '#' {
			d++
			continue
		}
		if d > 0 {
			d--
			continue
		}
		r += string(l)
	}
	return
}
