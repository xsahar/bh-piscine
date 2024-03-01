package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	r := os.Args[1:]

	if len(r) <= 0 {
		fmt.Println("File name missing")
	} else if len(r) > 1 {
		fmt.Println("Too many arguments")
	} else {
		fileinput, err := ioutil.ReadFile(r[0])
		if err != nil {
			fmt.Println("error in file reading")
		}
		fmt.Print(string(fileinput))
	}
}
