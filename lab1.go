package main

import (
	"fmt"
	"os"
	"bufio"
)
func code() {

}
func decode() {

}
func getInfo(st1,st2 string) string {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Print("input "+st1+"[default "+st2+"]:")
	language,_:=reader.ReadString('\n')
	if lng:=len(language); lng < 2 {
		return st2
	} else {
		return language[:lng-1]
	}
}
func main() {
	language := getInfo("language","ua")
	fmt.Println(language)
	fileName := getInfo("FileName","text")
	fmt.Println(fileName)
	resultName := getInfo("Result FileName","output")
	fmt.Println(resultName)
	direction := getInfo("direction(either 'code' or 'decode')","decode")
	fmt.Println(direction)
}
