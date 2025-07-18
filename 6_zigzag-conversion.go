func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	l := len(s)
	mod := (numRows - 1) * 2
	var ans string

	for j := 0; j < numRows; j++ {
		if j == 0 || j == numRows-1 {
			for i := 0; i < l; i++ {
				if i%mod == j {
					ans = ans + string(s[i])
				}
			}
		} else {
			for i := 0; i < l; i++ {
				if i%mod == j {
					ans = ans + string(s[i])
				}
				if i%mod == mod-j {
					ans = ans + string(s[i])
				}
			}
		}
	}

	return ans
}
