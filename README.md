# xkpassgo [![GoDoc](https://godoc.org/github.com/danmrichards/xkpassgo?status.svg)](https://godoc.org/github.com/danmrichards/xkpassgo) [![License](http://img.shields.io/badge/license-mit-blue.svg)](https://raw.githubusercontent.com/danmrichards/xkpassgo/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/danmrichards/xkpassgo)](https://goreportcard.com/report/github.com/danmrichards/xkpassgo)
A generator of [XKCD memorable passwords][1] implemented in Golang.

> Inspired by https://xkpasswd.net/s/ and https://github.com/bbusschots/hsxkpasswd

Word list based on the [10,000 most common][2] English words.

## Installation
If you are a Golang developer you can install the latest release like so:
```bash
$ go install github.com/danmrichards/xkpassgo/cmd/xkpassgo
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
      --word_length_max int             maximum word length (default 8)
      --word_length_min int             minimum word length (default 4)
```
> These options are heavily based on https://xkpasswd.net/s/

Configuration is determined in this order, each item takes precedence over the
item below it:

1. Flags
2. Configuration file
3. Defaults

### Configuration file
XKPassgo can load it's configuration from a JSON file. An example config file
looks like:
```json
{
 "num_words": 3,
 "word_length_min": 4,
 "word_length_max": 8,
 "case_transform": "ALTERNATE",
 "separator_character": "RANDOM",
 "separator_alphabet": [
  "!",
  "@",
  "$",
  "%",
  "^",
  "&",
  "*",
  "-",
  "_",
  "+",
  "=",
  ":",
  "|",
  "~",
  "?",
  "/",
  ".",
  ";"
 ],
 "padding_digits_before": 2,
 "padding_digits_after": 2,
 "padding_type": "FIXED",
 "padding_character": "RANDOM",
 "symbol_alphabet": [
  "!",
  "@",
  "$",
  "%",
  "^",
  "&",
  "*",
  "-",
  "_",
  "+",
  "=",
  ":",
  "|",
  "~",
  "?",
  "/",
  ".",
  ";"
 ],
 "padding_characters_before": 2,
 "padding_characters_after": 2,
 "random_increment": "AUTO"
}
```
> The config syntax is taken directly from https://xkpasswd.net/s/. It should be
possible to use a config file for both the binary and the website.

This file can exist anywhere and passed to the binary with the `--config` flag.

If you do not specify a `--config` flag the binary will look for a file called
`.xkpassgo.json` in your home directory.

If the given, or default, config file does not exist it will be ignored. The
configuration will then use the default values outlined in the usage above.

## TODO
- [ ] Subcommand to display config as JSON
- [ ] Subcommand to save config from flags
- [ ] Config presets from https://xkpasswd.net/s/

[1]: https://xkcd.com/936/
[2]: https://github.com/first20hours/google-10000-english