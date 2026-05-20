package main

import (
	"fmt"
	"log"
	"os"

	"github.com/danmrichards/xkpassgo/pkg/config"
	"github.com/danmrichards/xkpassgo/pkg/generator"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	if err := cfg.Validate(); err != nil {
		log.Fatal(err)
	}

	xkp := generator.NewXKPassword(cfg)

	pw, err := xkp.Generate()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(os.Stdout, pw)
}
