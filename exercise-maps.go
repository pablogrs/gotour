package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {

	string_map := make(map[string]int)
	splited_string := strings.Split(s, " ")

	for _, value:= range splited_string{
		string_map[value] += 1 
	}

	return string_map 
}

func main() {
	wc.Test(WordCount)
}
