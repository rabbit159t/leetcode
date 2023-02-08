func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	index := -1
	counter := make(map[byte]int)
	has_dup := false
	for i := 0; i < len(s); i++ {
		counter[s[i]]++
		if counter[s[i]] > 1 {
			has_dup = true
		}
		if s[i] != goal[i] {
			if index == -1 {
				index = i
			} else {
				new_s := []rune(s)
				new_s[i], new_s[index] = new_s[index], new_s[i]
				return string(new_s) == goal
			}
		}
	}
	if index == -1 {
		return has_dup
	} else {
		return false
	}
}
