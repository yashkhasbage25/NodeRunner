package utils

// InArray checks if a string str is present in array
func InArray(str string, arr []string) bool {

	for _, value := range arr {
		if value == str {
			return true
		}
	}
	return false
}

// StrSplit splits a string at the character c
func StrSplit(str string, c byte) []string {
	var i int
	for i = 1; i < len(str); i++ {
		if str[i] == c {
			break
		}
	}
	return []string{str[:i], str[i+1:]}
}
