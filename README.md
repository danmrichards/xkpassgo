# xkpassgo [![GoDoc](https://godoc.org/github.com/danmrichards/xkpassgo?status.svg)](https://godoc.org/github.com/danmrichards/xkpassgo) [![License](http://img.shields.io/badge/license-mit-blue.svg)](https://raw.githubusercontent.com/danmrichards/xkpassgo/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/danmrichards/xkpassgo)](https://goreportcard.com/report/github.com/danmrichards/xkpassgo) ![](https://github.com/danmrichards/xkpassgo/workflows/test/badge.svg)
A generator of [XKCD memorable passwords][1] implemented in Golang.

> Inspired by https://xkpasswd.net/s/ and https://github.com/bbusschots/hsxkpasswd

Word list based on the [10,000 most common][2] English words.

## Installation
If you are a Golang developer you can install the latest release like so:
```bash
$ go get github.com/icanhazpython/xkpassgo/cmd/xkpassgo
```
Otherwise you can download a binary for your platform from the releases page.

## Usage
To just generate a new password:
```bash
$ xkpassgo
```

The binary has several configuration options that can be set:
```bash
Usage of xkpassgo:
      --case_transform string           case transformation, allowed values: LOWER, UPPER, RANDOM, NONE, ALTERNATE, CAPITALISE, INVERT (default "ALTERNATE")
      --config string                   path to config file (default "/home/dan/.xkpassgo.json")
      --num_words int                   number of words (default 3)
      --pad_to_length int               length to pad the password to, will be ignored if less than the generated password length (default 8)
      --padding_alphabet strings        comma-separated list of characters to pad the password with (default [!,@,$,%,^,&,*,-,_,+,=,:,|,~,?,/,.,;])
      --padding_character string        character to pad the password with (default "RANDOM")
      --padding_characters_after int    number of characters to pad before the password (default 2)
      --padding_characters_before int   number of characters to pad before the password (default 2)
      --padding_digits_after int        number of digits to pad before the password (default 2)
      --padding_digits_before int       number of digits to pad before the password (default 2)
      --padding_type string             padding type, allowed values: FIXED, ADAPTIVE (default "FIXED")
      --separator_alphabet strings      comma-separated list of characters to separate password parts (default [!,@,$,%,^,&,*,-,_,+,=,:,|,~,?,/,.,;])
      --separator_character string      character to separate password parts (default "RANDOM")
      --separator_characters int        number of characters to separated password parts with (default 1)
      --word_length_max int             maximum word length (default 8)
      --word_length_min int             minimum word length (default 4)
```
> These options are heavily based on https://xkpasswd.net/s/


[1]: https://xkcd.com/936/
[2]: https://github.com/first20hours/google-10000-english
