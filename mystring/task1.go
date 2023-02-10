package mystring

func Reverse(str string) string {
	strRune := []rune(str)
	for i := 0; i < len(strRune)/2; i++ {
		var r1 rune
		r1 = strRune[i]
		strRune[i] = strRune[len(strRune)-1-i]
		strRune[len(strRune)-1-i] = r1
	}
	return string(strRune)
}
