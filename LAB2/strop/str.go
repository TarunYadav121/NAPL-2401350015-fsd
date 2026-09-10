package strop

// func PrintStr(s string) string {
//     return s
// }

func Reverse(s string)string {
	n := len(s)
	result := ""

	for i := n - 1; i >= 0; i-- {
		result += string(s[i])
	}

	return result
}


func CountVowels(s string) int {
	count := 0

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
			ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
			count++
		}
	}

	return count
}

