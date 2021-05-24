package main

import (
	"fmt"
	"strings"
)

func main() {
	numbers := "Lorem ipsum dolor sit amet"
	num, occurrence := countWords(numbers)
	fmt.Println(num, occurrence) //5map[amet:1dolor:1ipsum:1lorem:1sit:1]
}

func countWords(sentence string) (int, map[string]int) {
	//tijelo funkcije
	resultMap := map[string]int{}
	tokens := strings.Split(sentence, " ")

	for _, value := range tokens {
		withoutDot := strings.Trim(value, ".")
		withoutComma := strings.Trim(withoutDot, ",")
		word := strings.ToLower(withoutComma)

		value, ok := resultMap[word]
		if ok {
			resultMap[word] = value + 1
		} else {
			resultMap[word] = 1
		}

	}

	return len(resultMap), resultMap
}
