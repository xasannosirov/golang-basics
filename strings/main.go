package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	fmt.Println(
		strings.Contains("test", "es"),
		strings.Count("test", "t"),
		strings.HasPrefix("test", "te"),
		strings.HasSuffix("test", "st"),
		strings.Index("test", "e"),
		strings.Join([]string{"hello", "world"}, "-"),
		strings.Repeat("a", 5),
		strings.Replace("blanotblanot", "not", "***", -1),
		strings.Split("a-b-c-d-e", "-"),
		strings.ToLower("TEST"),
		strings.ToUpper("test"),
		strings.Trim("tetstet", "te"),
	)

	fmt.Println(unicode.IsDigit('1'))
	fmt.Println(unicode.IsLetter('a'))
	fmt.Println(unicode.IsLower('A'))
	fmt.Println(unicode.IsUpper('A'))
	fmt.Println(unicode.IsSpace('\t'))
	fmt.Println(unicode.Is(unicode.Latin, 'ы'))
	fmt.Println(string(unicode.ToLower('F')))
	fmt.Println(string(unicode.ToUpper('f')))

	var en = "english"
	var ru = "русский"
	fmt.Println(len(en), len(ru))
	fmt.Println(utf8.RuneCountInString(en), utf8.RuneCountInString(ru))
}
