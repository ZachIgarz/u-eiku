package ia

import (
	"bufio"
	"log"
	"os"
)

//KeyWords is the global struct of keywords
type KeyWords struct {
	Words []string
}

//NewKeyWords is the initializatos of keywords
func NewKeyWords() *KeyWords {
	return &KeyWords{
		Words: readWords(),
	}
}

func readWords() []string {
	var words []string
	file, err := os.Open("/home/zach/PruebasZach/u-eiku/ia/words.txt")
	if err != nil {

		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words
}
