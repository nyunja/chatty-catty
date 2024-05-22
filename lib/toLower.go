package lib

func toLower(s string) string {
	result := ""
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			result += string(ch)
		} else {
			result += string(ch + 32)
		}
	}
	return result
}
