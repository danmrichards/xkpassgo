package main

import (
	"fmt"
	"os"

	"github.com/danmrichards/xkpassgo/pkg/config"
	"github.com/danmrichards/xkpassgo/pkg/generator"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		exitErr(err)
	}

	if err = cfg.Validate(); err != nil {
		exitErr(err)
	}

	xkp := generator.NewXKPassword(cfg)

	pw, err := xkp.Generate()
	if err != nil {
		exitErr(err)
	}

	fmt.Println(pw)
}

func exitErr(err error) {
	fmt.Println(err)
	os.Exit(1)
}
