package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println("## Soal 1 ##")
	//1.a
	stringOne := "19374048"
	//1.b
	stringTwo := "aiueobcd"
	//1.c
	fmt.Println("find 7 in string is on index :", FindSevenInString(stringOne))
	//1.d --> This is the function() that can recieve 2 params string
	PrintTheRestStringAfterSeven(stringOne, stringTwo)
}

func FindSevenInString(s string) int{
	sArr := strings.Split(s, "")
	for index, row := range sArr {
		if row == "7" {
			return index
		}
	}
	return 0
}

//PrintTheRestStringAfterSeven() is the function that can recieve 2 params string
func PrintTheRestStringAfterSeven(sOne, sTwo string){
	sArr := strings.Split(sTwo, "")
	fmt.Print("print rest of string after index of seven on previous string('"+sOne+"') : ")
	for i := FindSevenInString(sOne); i < len(sArr); i++{
		fmt.Print(sArr[i])
	}
	fmt.Print("\n\n")
}
