package merge_strings_alternately

func mergeAlternately(word1, word2 string) string {
	var byteresult []byte
	i := 0

	for i < len(word1) && i < len(word2) {
		byteresult = append(byteresult, word1[i])
		byteresult = append(byteresult, word2[i])
		i++
	}

	result := string(byteresult)

	if i >= len(word1) && i < len(word2) {
		result += string(word2[i:])
	}

	if i >= len(word2) && i < len(word1) {
		result += string(word1[i:])
	}

	return result
}
