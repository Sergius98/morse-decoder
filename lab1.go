package main

import (
	"fmt"
	"os"
	"bufio"
	"io/ioutil"
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
func check(e error) {
    if e != nil {
        panic(e)
    }
}
func getData(file string)string{
	dat, err := ioutil.ReadFile(file)
	check(err)
	return string(dat)
}
func getRunes(file string)[]rune{
	str := getData(file)
	return []rune(str)
}
func getWord(smb,str string, i int)(int, string){
	word:=""
	for pos:=i; pos<len(str); pos++{
		if str[pos]==smb[0]{
			return pos+1, word
		} else {
			word+=str[pos:pos+1]
		}
	}
	return len(str), word
}
func updateDictionary(code,key string, reverse bool, dict map[string]string) map[string]string{
	if reverse {
		dict[code] = key
	} else {
		dict[key] = code
	}
	return dict
}
func getDict(file string, reverse bool)map[string]string{
	str, dict, key, code := getData(file),make(map[string]string),"",""
	for i:=0; i<len(str);{
		i,key = getWord(" ",str,i)
		i,code = getWord("\n",str,i)
		if i >= len(str) {
			return dict
		}
		dict = updateDictionary(code, key, reverse, dict)
	}
	return dict
}
func root(language string, fileName string, resultName string, reverse bool) {
	arr := getRunes(fileName)
	fmt.Println(arr)
	dict := getDict(language, reverse)
	fmt.Println(dict)
}
func getVars() (string, string, string, string) {
	language := getInfo("language","ua")
	fileName := getInfo("FileName","part1")
	resultName := getInfo("Result FileName","output")
	direction := getInfo("direction(either 'code' or 'decode')","decode")
	return language, fileName, resultName, direction
}
func main() {
	language, fileName, resultName, direction := getVars()
	if direction == "code" {
		root(language, fileName, resultName, false)
	} else if direction == "decode" {
		root(language, fileName, resultName, true)
	} else {
		fmt.Println("wrong direction")
	}
}
