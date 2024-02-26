package validanagram

func validAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	res := make(map[rune]int, len(s))
	for _, char := range s {
		res[char]++
	}

	for _, char := range t {
		res[char]--
	}

	for _, v := range res {
		if v != 0 {
			return false
		}
	}

	return true
}
