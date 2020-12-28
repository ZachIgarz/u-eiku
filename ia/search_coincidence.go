package ia

import (
	"fmt"
	"strings"
)

//tablas de hash

//WordRepeticions ...
func WordRepeticions(text string) {
	var diccionary map[string]int

	separatedWords := strings.Split(text, " ")

	for _, word := range separatedWords {

		//TODO :existe algo como in de js en go ?? D: deam it
		ok := if normalize(word) in diccionary {
			++dicdiccionary[normalize(word)]
		}
		fmt.Println(word)
	}

}

func normalize(word string) string {

}

func searhInDiccionary(word string, diccionary  map[string]int) bool{

	var count word 
	for _, wordInDicc := range diccionary{
		if word == wordInDicc {

		}
	}


}