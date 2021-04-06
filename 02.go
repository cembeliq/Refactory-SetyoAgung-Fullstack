package main

import (
	"fmt"
	"regexp"
)

var replaceWord [7]string
var text string
var x int = -1

func main() {
	
	replaceWord = [...]string{"dolor", "elit", "quis", "nisi", "fugiat", "proident", "laborum"}
	
	text = `Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
	Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
	Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.
	Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
	`
	output := replace(text)
	fmt.Println(output)
}

func process(word string) string{
	filterVal := fmt.Sprintf(`\b%s\b`, word)

	var regex, _ = regexp.Compile(filterVal)

	text = regex.ReplaceAllString(text, generateAsterix(word))

	return text
	
}

func replace(text string) string {
	x = x + 1
	if x>len(replaceWord)-1 {
		return text
	}
	return replace(process(replaceWord[x]))
}

func generateAsterix(word string) string{
	countChar := len([]rune(word))
	asterix := ""

	for i:=0;i<countChar;i++{
		asterix += "*"
	}

	return asterix

}