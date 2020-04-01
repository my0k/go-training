package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// tryDiscard()
	tryReadAll()
}

func tryDiscard() {
	discard := ioutil.Discard
	num, err := discard.Write([]byte("this string is discarded"))
	if err != nil {
		fmt.Printf("got err: %v\n", err)
	}

	fmt.Printf("%d bytes written.\n", num)
}

func tryReadAll() {
	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")

	for i := 0; i < 3; i++ {
		b, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", b)
	}
}
