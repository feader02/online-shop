package utils

func CreateTrigrams(s string) []string {
	var trigrams []string
	for i := 0; i < len(s)-2; i++ {
		trigrams = append(trigrams, s[i:i+3])
	}
	if len(s) < 3 {
		trigrams = append(trigrams, s)
	}
	return trigrams
}
