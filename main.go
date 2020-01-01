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

	var hex string
	for _, letter := range text {
		a, _ := strconv.ParseInt(fmt.Sprintf("%d", letter), 10, 32)
		u := fmt.Sprintf("%X", letter)

		if a > 127 {
			hex += "&#x" + string(u)
		} else {
			hex += string(letter)
		}
	}
	fmt.Println(hex)
}
