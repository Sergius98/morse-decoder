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

	f, err := os.Create("input_long")
  check(err)
	defer f.Close()
	for i:=0; i<19; i++ {
		f.Write(dat)
	}
	f.Sync()
	fmt.Println("input successfully finished")

	dat2, err2 := ioutil.ReadFile("output")
	check(err2)

	f2, err2 := os.Create("output_long")
  check(err2)
	defer f2.Close()
	for i:=0; i<19; i++ {
		f2.Write(dat2)
	}
	f2.Sync()
	fmt.Println("output successfully finished")

}
