package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	var filename string
	flag.StringVar(&filename, "file", "", "file name")
	flag.StringVar(&filename, "f", "", "file name")

	var charcount bool
	flag.BoolVar(&charcount, "chars", false, "char count")
	flag.BoolVar(&charcount, "c", false, "char count")

	var linecount bool
	flag.BoolVar(&linecount, "lines", false, "line count")
	flag.BoolVar(&linecount, "l", false, "line count")

	var wordcount bool
	flag.BoolVar(&wordcount, "word", false, "word count")
	flag.BoolVar(&wordcount, "w", false, "word count")

	flag.Parse()

	if filename == "" {
		flag.Usage()
	} else {
		contents, err := ioutil.ReadFile(filename)
		check(err)

		s := string(contents)
		length := len(s)
		count := strings.Count(s, "\n")

		//s1 := strings.Replace(s, "\r\n", " ", count-1)

		split := strings.Fields(s)

		if !charcount && !linecount && !wordcount {
			fmt.Printf("%d %d %d %s", count, len(split), length, filename)
		} else if charcount {
			fmt.Printf("%d %s", length, filename)
		} else if linecount {
			fmt.Printf("%d %s", count, filename)
		} else if wordcount {
			fmt.Printf("%d %s", len(split), filename)
		}
	}
}
