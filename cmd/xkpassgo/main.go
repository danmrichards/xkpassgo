package main

import (
	"fmt"
	"log"

	"github.com/danmrichards/xkpassgo/internal/config"
	"github.com/danmrichards/xkpassgo/internal/generator"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Validate config.

	xkp := generator.NewXKPassword(cfg)

	pw, err := xkp.Generate()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pw)
}
