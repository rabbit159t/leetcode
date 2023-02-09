func areAlmostEqual(s1 string, s2 string) bool {
	index := -1
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if index == -1 {
				index = i
			} else {
				// Do swap
				new_s := []rune(s1)
				new_s[index], new_s[i] = new_s[i], new_s[index]
				return string(new_s) == s2
			}
		}
	}
	if index != -1 {
		return false
	}
	return true
}