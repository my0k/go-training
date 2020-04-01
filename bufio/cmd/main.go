package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func lineSample1() {
	r := bufio.NewScanner(strings.NewReader("aaa bbb ccc"))
	r.Buffer(make([]byte, 1024), int(1<<30))
	r.Split(bufio.ScanLines)

	r.Scan()
	fmt.Printf("Input: %s\n", r.Text())
	fmt.Printf("Input: %s\n", r.Text())
}

func wordSample1() {
	r := bufio.NewScanner(strings.NewReader("aaa bbb ccc"))
	r.Buffer(make([]byte, 1024), int(1<<30))
	r.Split(bufio.ScanWords)

	r.Scan()
	fmt.Printf("Input: %s\n", r.Text())
	fmt.Printf("Input: %s\n", r.Text())
	r.Scan()
	fmt.Printf("Input: %s\n", r.Text())
	fmt.Printf("Input: %s\n", r.Text())
	r.Scan()
	fmt.Printf("Input: %s\n", r.Text())
	fmt.Printf("Input: %s\n", r.Text())
}

func initBySpecialString() {
	r := bufio.NewScanner(strings.NewReader("aaa bbb ccc"))
	r.Buffer([]byte("VimVimVim"), int(1<<30))

	fmt.Printf("Input: %s\n", r.Text()) // return ""
}

func bufWriterSample1() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprintf(w, "Hello, ")
	fmt.Fprintf(w, "world!\n")
	time.Sleep(time.Second * 2)
	w.Flush()
}

func main() {
	// lineSample1()
	// initBySpecialString()
	// wordSample1()

	bufWriterSample1()
}
