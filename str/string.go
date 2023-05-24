package str


func SubString(str string, start int, end int) string {
	if len(str) < start {
		return ""
	}
	if len(str) < end {
		return str[start:]
	}
	chars := []rune(str)
	return string(chars[start:end])
}