package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("No input text. Please specify the text to html encode.")
		return
	}

	text := os.Args[1]

	var encoded string
	for _, letter := range text {
		a, _ := strconv.ParseInt(fmt.Sprintf("%d", letter), 10, 32)
		u := fmt.Sprintf("%d", letter)

		if a > 127 {
			encoded += "&#" + string(u)
		} else {
			encoded += string(letter)
		}
	}
	fmt.Println(encoded)
}
