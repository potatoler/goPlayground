package postring

// split a string at indexed place, returning the two strings
func Split(str string, index int) (string, string) {
	rawStr := []byte(str)
	return string(rawStr[:index]), string(rawStr[index:])
}

// character by character reverse a string
func Reverse(str string) string {
	rawString := []byte(str)
	length := len(rawString)
	for index := 0; index < length/2; index++ {
		rawString[index], rawString[length-index-1] = rawString[length-index-1], rawString[index]
	}
	return string(rawString)
}
