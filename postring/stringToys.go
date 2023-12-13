package postring

// so-called "Uniqualize" can traverse a string, and copy characters that differ from the very last one to a new string
func Uniqualize(str string) string {
	rawString := []byte(str)
	length := len(rawString)
	var rawAnswer []byte
	if length == 0 || length == 1 {
		return str
	}
	rawAnswer = append(rawAnswer, rawString[0])
	for i := 1; i < length; i++ {
		if rawString[i] != rawString[i-1] {
			rawAnswer = append(rawAnswer, rawString[i])
		}
	}
	return string(rawAnswer)
}
