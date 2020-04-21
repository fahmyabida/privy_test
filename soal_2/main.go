package main

import (
	"fmt"
	"regexp"
)

func main(){
	fmt.Println("## Soal 2 ##")
	stringThreePointOne   := "asfjkah[fsaf]"
	stringThreePointTwo   := "[  ]asgasda"
	stringThreePointThree := "asfsdasdasd["
	fmt.Println(stringThreePointOne, "\t:", DetectBracket(stringThreePointOne))
	fmt.Println(stringThreePointTwo, "\t:", DetectBracket(stringThreePointTwo))
	fmt.Println(stringThreePointThree, "\t:", DetectBracket(stringThreePointThree))
	stringFourPointOne	 := "{safa}asf"
	stringFourPointTwo	 := "{}[]"
	stringFourPointThree := "{asfaf]"
	fmt.Println(stringFourPointOne, "\t:", DetectBracket(stringFourPointOne))
	fmt.Println(stringFourPointTwo, "\t\t:", DetectBracket(stringFourPointTwo))
	fmt.Println(stringFourPointThree, "\t:", DetectBracket(stringFourPointThree))
}

func DetectBracket(s string) bool{
	regexExp := regexp.MustCompile(`\[(.*)\]|\{(.*)\}`)
	return regexExp.MatchString(s)
}
