package main

import (
	"fmt"
	"regexp"
)

func main() {
	pattern := [...]string{"aku", "ingin", "dapat"}
	
	text := `Aku ingin begini
	Aku ingin begitu
	Ingin ini itu banyak sekali
	
	Semua semua semua
	Dapat dikabulkan
	Dapat dikabulkan
	Dengan kantong ajaib
	
	Aku ingin terbang bebas
	Di angkasa
	Hei… baling baling bambu
	
	La... la... la...
	Aku sayang sekali
	Doraemon
	
	La... la... la...
	Aku sayang sekali
	`

	for i:=0; i<len(pattern);i++ {
		for index, element := range countWord(pattern[i], text) {
			fmt.Println(index, "=>", element)
		}
	}
}

func countWord(word string, text string) map[string]int {
	counts := make(map[string]int)

	filterVal := fmt.Sprintf(`(?i)%s\b`, word)
	var regex, _ = regexp.Compile(filterVal)


	var str = regex.FindAllString(text, -1)
	counts[word] = len(str)
	return counts

}