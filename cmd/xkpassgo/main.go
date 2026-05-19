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

	err = cfg.Validate()
	if err != nil {
		exitErr(err)
	}

	xkp := generator.NewXKPassword(cfg)

	pw, err := xkp.Generate()
	if err != nil {
		exitErr(err)
	}

	fmt.Fprintln(os.Stdout, pw)
}

func exitErr(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
