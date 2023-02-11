func equalFrequency(word string) bool {
	count := make(map[rune]int)
	for _, c := range word {
		count[c]++
	}
	if len(count) == 1 {
		return true
	}

	freq := make(map[int]int)
	freqSlice := make([][]int, 0)
	for _, v := range count {
		freq[v]++
	}

	for k, v := range freq {
		freqSlice = append(freqSlice, []int{k, v})
	}
	if len(freqSlice) == 1 {
		return freqSlice[0][0] == 1
	} else if len(freq) != 2 {
		return false
	}
	if freqSlice[0][0] == 1 && freqSlice[0][1] == 1 {
		return true
	} else if freqSlice[1][0] == 1 && freqSlice[1][1] == 1 {
		return true
	} else {
		if freqSlice[0][0]-freqSlice[1][0] == 1 {
			return freqSlice[0][1] == 1
		} else if freqSlice[1][0]-freqSlice[0][0] == 1 {
			return freqSlice[1][1] == 1
		}
	}
	return false
}