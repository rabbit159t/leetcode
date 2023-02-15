func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	str := make([]byte, 0)
	n := len(s)
	k := numRows<<1 - 2
	for i := 0; i < numRows; i++ {
		j := i
		first := true
		step1, step2 := k-j*2, j*2
		if step1 == 0 {
			step1 = k
		}
		if step2 == 0 {
			step2 = k
		}

		for j < n {
			str = append(str, s[j])
			if first {
				j += step1
				first = false
			} else {
				j += step2
				first = true
			}
		}
	}
	return string(str)
}