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
		if normalize(word) in diccionary {
			++dicdiccionary[normalize(word)]
		}
		fmt.Println(word)
	}

}

func normalize(word string) string {

}
