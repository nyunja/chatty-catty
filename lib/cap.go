package lib

func Cap(s string) string {
	result := ""
	if string(s[0]) >= "a" && string(s[0]) <= "z" {
		result += string(s[0] - 32)
	}
	return result + s[1:]
}