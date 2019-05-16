package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mouuii/go-src/io/myreader"
)

func main1() {
	reader := strings.NewReader("Clear is better than clever")
	p := make([]byte, 4)

	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println(string(p[:n]))
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(string(p[:n]))
	}
}
func main2() {
	reader := myreader.NewAlphaReader("Hello! It's 9am, where is the sun?")
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:n]))
	}
	fmt.Println()
}
func main3() {
	//chaining readers
	//The standard library has many readers already implemented. It is a common idiom to use a reader as the source of another reader. This chaining of readers allows one reader to reuse logic from another as is done in the following source snippet which updates the alphaReader to accept an io.Reader as its source. This reduces the complexity of the code by pushing stream housekeeping concerns to the root reader.

	reader := myreader.NewAlphaChainReader(strings.NewReader("adsfafa asdfadf!@!df asd"))
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:n]))
	}
	fmt.Println()
}
func main() {
	proverbs := []string{
		"channels orchaestrate",
		"sdafadsdfasdfasdf",
		"asdfafasf",
		"dont panic",
	}
	var write bytes.Buffer

	for _, p := range proverbs {
		n, err := write.Write([]byte(p))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(p) {
			fmt.Println("fail to write data")
			os.Exit(1)
		}
	}
	fmt.Println(write.String())
}
