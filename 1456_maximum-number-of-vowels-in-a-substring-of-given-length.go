func maxVowels(s string, k int) int {
	max := 0
	count := 0

	for i := 0; i < len(s)-k+1; i++ {
		if i == 0 {
			substr := string(s[0 : k])

			count = countVowel(substr)
			if count > max {
				max = count
			}
            continue
		}

        if isVowel(s[i-1]) {
            count--
        }

        if isVowel(s[i+k-1]) {
            count++
        }

        if count > max {
			max = count
		}
	}

	return max
}

func countVowel(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' ||
			s[i] == 'e' ||
			s[i] == 'i' ||
			s[i] == 'o' ||
			s[i] == 'u' {
			count++
		}
	}

	return count
}

func isVowel(b byte) bool {
	return b == 'a' ||
		b == 'e' ||
		b == 'i' ||
		b == 'o' ||
		b == 'u'
}