package kata

import "strings"

func AbbrevName(name string) string {
	words := strings.Split(name, " ")

	var initials []string
	for _, word := range words {
		first := strings.ToUpper(string([]rune(word)[0]))
		initials = append(initials, first)
	}

	return strings.Join(initials, ".")
}

//https://www.codewars.com/kata/57eadb7ecd143f4c9c0000a3/solutions/go
