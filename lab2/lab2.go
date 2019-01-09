package main

import (
	"fmt"
	"os"
	//"bufio"
	"io/ioutil"
	"strings"
)
func encode(arr []rune, dict map[string]string) []string {
	str := make([]string, len(arr))
	for i:=0;i<len(arr);i++{
		str[i] = dict[string(arr[i])] + " ";
	}
	return str
}
func getCode(arr []rune, i int)(int, string){
	word := []rune("")
	for pos:=i; pos<len(arr); pos++{
		if arr[pos]==32{
			return pos+1, string(word)
		} else {
			word = append(word, arr[pos])
		}
	}
	return len(arr), string(word)
}
func decode(arr []rune, dict map[string]string) []string {
	retval := make([]string, len(arr))
	code := ""
	j:=0;
	for i:=0;i<len(arr);{
		i,code = getCode(arr,i)
		retval[j] = dict[code]
		j++
	}
	return retval[:j]
}
// func getInfo(st1,st2 string) string {
// 	var reader = bufio.NewReader(os.Stdin)
// 	fmt.Print("input "+st1+"[default "+st2+"]:")
// 	language,_:=reader.ReadString('\n')
// 	if lng:=len(language); lng < 2 {
// 		return st2
// 	} else {
// 		return language[:lng-1]
// 	}
// }
func check(e error) {
    if e != nil {
        panic(e)
    }
}
func getData(file string)string{
	dat, err := ioutil.ReadFile(file)
	check(err)
	return strings.ToLower(string(dat))
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
	if (code[0:1]==" ")&&(key==""){
		code,key = code[1:]," "
	}
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
	arr, dict:= getRunes(fileName), getDict(language, reverse)
	var str []string
	if reverse {
		str = decode(arr, dict)
	} else {
		str = encode(arr, dict)
	}
	f, err := os.Create(resultName)
  check(err)
	defer f.Close()
	for i:=0; i<len(str); i++ {
		f.WriteString(str[i])
		//err := ioutil.WriteFile(resultName, []byte(str), 0644)// Unix permission bits
	}
	f.Sync()
	fmt.Println("program successfully finished")
}
//func getVars() (string, string, string, string) {
	// language := getInfo("language","ua")
	// fileName := getInfo("FileName","input")
	// resultName := getInfo("Result FileName","output")
	// direction := getInfo("direction(either 'encode' or 'decode')","decode")
	// return language, fileName, resultName, direction
//}
 func main() {
 	root("ua", "out1", "kke", false)

}
