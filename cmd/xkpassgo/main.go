package main

import (
	"fmt"
	"log"

	"github.com/danmrichards/xkpassgo/internal/generator"
)

const (
	numWords   = 3
	minWordLen = 4
	maxWordLen = 8
)

func main() {
	// TODO: Config parsing from flags and/or JSON.

	xkp := generator.NewXKPassword(&generator.Config{
		NumWords:   numWords,
		WordLenMin: minWordLen,
		WordLenMax: maxWordLen,
	})

	pw, err := xkp.Generate()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pw)
}
