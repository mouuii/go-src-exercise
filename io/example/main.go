package main

import (
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
func main() {
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
