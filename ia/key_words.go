package ia

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//KeyWords is the globar variable of keywords
var KeyWords = make([]string, 0)

//Leer ..
func Leer() {
	file, err := os.Open("/home/zach/PruebasZach/u-eiku/ia/words.txt")
	if err != nil {

		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		KeyWords = append(KeyWords, scanner.Text())
	}

	fmt.Println(KeyWords[0], KeyWords[2])
}
