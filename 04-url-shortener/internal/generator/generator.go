package generator

const countOfAlphabet = 52

func GenerateURL(num int) string {
	reverseGeneratedURL := ""

	for num > 0 {
		alphabetIndex := num % countOfAlphabet
		if alphabetIndex == 0 {
			alphabetIndex = countOfAlphabet
			num = num / (countOfAlphabet + 1)
		} else {
			num /= countOfAlphabet
		}

		if alphabetIndex > countOfAlphabet/2 {
			reverseGeneratedURL += string(rune('A' - 1 + alphabetIndex - (countOfAlphabet / 2)))
		} else {
			reverseGeneratedURL += string(rune('a' - 1 + alphabetIndex))
		}
	}

	shortURL := ""
	for i := len(reverseGeneratedURL) - 1; 0 <= i; i-- {
		shortURL += string(reverseGeneratedURL[i])
	}
	return shortURL
}
