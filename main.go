package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	f, err := os.Open("C:\\")
	if err != nil {
		log.Fatal(err)
	}
	l, err := f.Readdir(0)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range l {
		fmt.Printf("%v\n", f.Name())
	}
}
