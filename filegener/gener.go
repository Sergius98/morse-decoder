package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}
func main() {
	dat, err := ioutil.ReadFile("input")
	check(err)
	dat12, err := ioutil.ReadFile("input_very_short")
	check(err)

	f, err := os.Create("input_long")
  check(err)
	defer f.Close()
	for i:=0; i<4; i++ {
		f.Write(dat)
	}
	f.Write(dat12)
	f.Write(dat12)
	f.Sync()
	fmt.Println("input successfully finished")

	dat2, err := ioutil.ReadFile("output")
	check(err)
	dat22, err := ioutil.ReadFile("output_very_short")
	check(err)

	f2, err := os.Create("output_long")
  check(err)
	defer f2.Close()
	for i:=0; i<4; i++ {
		f2.Write(dat2)
	}
	f2.Write(dat22)
	f2.Write(dat22)
	f2.Sync()
	fmt.Println("output successfully finished")

}
