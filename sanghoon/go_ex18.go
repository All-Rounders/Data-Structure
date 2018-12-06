package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("Yuin", "Yuin Lee")
	fmt.Println(match)
	match, _ = regexp.MatchString("Yuin Lee", "Lee")
	fmt.Println(match)

	parse, err := regexp.Compile("[Yy]uin")

	if err != nil {
		fmt.Printf("Error compiling RE: %s\n", err)
	} else {
		fmt.Println(parse.MatchString("Yuin Lee"))
		fmt.Println(parse.MatchString("yuin Lee"))
		fmt.Println(parse.MatchString("Y uin Lee"))
		fmt.Println(parse.ReplaceAllString("yuin Yuin", "YUIN"))
	}
}
