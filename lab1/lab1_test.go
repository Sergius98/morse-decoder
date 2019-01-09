package main

import (
	"fmt"
	"testing"
)

 func TestGetRunes(t *testing.T){

	 p1 := []rune(getRunes("get_vars"))
 	 p2 := [5]rune{45, 45, 32, 46,45}

	 for i:=0;i<5;i++{
 	 	if p1[i] != p2[i] {
			t.Error("problem with runes")
 	 	}
	 }
 }

func TestGetDict(t *testing.T) {
	p := getDict("short_ua",true)
	//fmt.Println(p,"\n")
	if p["....-"] != "4" || p["...-"] != "ж" ||  p["...--"] != "з"{
		t.Error("Ploblem with GetDict expected")
	 }
}

func TestGetData(t *testing.T){
	p := getData("short_ua")
	//fmt.Println(len(p),"\n")
	if p != "4 ....-\nж ...-\nз ...--\nend" {
		t.Error("Ploblem with GetData expected")
	 }
}

func TestGetWord(t *testing.T){
	str, key, code:= "4 ....-\n", "",""
	for i:=0; i<len(str);{
		i,key = getWord(" ",str,i)
			if key != "4"{
			t.Error("Ploblem with GetWord(key) expected")
			//fmt.Println(i,key,"!\n")
		}
		i,code = getWord("\n",str,i)
					if code != "....-"{
			t.Error("Ploblem with GetWord(code) expected")
			//fmt.Println(i,code,"&\n")
		}
		if i >= len(str) {
			break
		}
	}
}


func TestFoo(t *testing.T) {
	t.Run("A=1", func (t *testing.T){
		root("ua", "output", "input", false)
	})
  t.Run("A=2", func (t *testing.T){
		root("ua", "input", "output", true)
	})

	t.Run("A=4", func (t *testing.T){
		str:= encode(getRunes("out1"),getDict("ua",false))
		fmt.Printf("%s",str)

 		 if str != "-- -.-- -.- --- .-.. .- "{
 			t.Error("True Decode Test crash")
 		}
	})
		t.Run("A=3", func (t *testing.T){
	 		if decode(getRunes("rtest")[0:22],getDict("ua",true)) != "маруся"{
	 			t.Error("Error decode Test")
	 		}
		})

}
