package find_the_difference

func findTheDifference(s string, t string) byte {
	result := 0

	for i := 0; i < len(s); i++ {
		result += int(t[i]) - int(s[i])
	}

	result += int(t[len(s)])

	return byte(result)
}
