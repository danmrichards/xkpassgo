package config

import (
	"errors"

	"github.com/jessevdk/go-flags"
)

var (

	// DefaultAlphabet is the default character set used for the separator
	// and padding modifications.
	DefaultAlphabet = []string{
		"!", "@", "$", "%", "^", "&", "*", "-", "_",
		"+", "=", ":", "|", "~", "?", "/", ".", ";",
	}
)

// GeneratorConfig represents the configuration for the password generator.
type GeneratorConfig struct {
	NumWords                int      `long:"num_words" default:"3" description:"number of words"`
	WordLenMin              int      `long:"word_length_min" default:"4" description:"minimum word length"`
	WordLenMax              int      `long:"word_length_max" default:"8" description:"maximum word length"`
	CaseTransform           string   `long:"case_transform" default:"ALTERNATE" choice:"LOWER" choice:"UPPER" choice:"RANDOM" choice:"NONE" choice:"ALTERNATE" choice:"CAPITALISE" choice:"INVERT" description:"case transformation"`
	SeparatorCharacters     int      `long:"separator_characters" default:"1" description:"number of characters to use to separate words"`
	SeparatorCharacter      string   `long:"separator_character" default:"RANDOM" description:"character to separate password parts"`
	SeparatorAlphabet       []string `long:"separator_alphabet" default:"!" default:"@" default:"$" default:"%" default:"^" default:"&" default:"*" default:"-" default:"_" default:"+" default:"=" default:":" default:"|" default:"~" default:"?" default:"/" default:"." default:";" description:"comma-separated list of characters to separate password parts"`
	PaddingDigitsBefore     int      `long:"padding_digits_before" default:"2" description:"number of digits to pad before the password"`
	PaddingDigitsAfter      int      `long:"padding_digits_after" default:"2" description:"number of digits to pad after the password"`
	PaddingType             string   `long:"padding_type" default:"FIXED" choice:"FIXED" choice:"ADAPTIVE" description:"padding type"`
	PaddingCharacter        string   `long:"padding_character" default:"RANDOM" description:"character to pad the password with"`
	SymbolAlphabet          []string `long:"symbol_alphabet" default:"!" default:"@" default:"$" default:"%" default:"^" default:"&" default:"*" default:"-" default:"_" default:"+" default:"=" default:":" default:"|" default:"~" default:"?" default:"/" default:"." default:";" description:"comma-separated list of characters to pad the password with"`
	PadToLength             int      `long:"pad_to_length" default:"8" description:"length to pad the password to, will be ignored if less than the generated password length"`
	PaddingCharactersBefore int      `long:"padding_characters_before" default:"2" description:"number of characters to pad before the password"`
	PaddingCharactersAfter  int      `long:"padding_characters_after" default:"2" description:"number of characters to pad before the password"`
}

func Load() (gc *GeneratorConfig, err error) {
	// Parse args
	var gcReal GeneratorConfig
	var parser *flags.Parser = flags.NewParser(&gcReal, flags.Default)
	if _, err = parser.Parse(); err == nil {
		gc = &gcReal
	}
	return
}

// Validate returns an error if the current configuration is not valid.
func (gc *GeneratorConfig) Validate() error {
	if gc.NumWords < 0 {
		return errors.New("num_words must be a positive integer")
	}

	if gc.WordLenMin < 0 {
		return errors.New("word_len_min must be a positive integer")
	}
	if gc.WordLenMax < 0 {
		return errors.New("word_len_max must be a positive integer")
	}
	if gc.WordLenMax < gc.WordLenMin {
		return errors.New("word_len_max cannot be less than word_len_min")
	}

	if gc.PaddingDigitsBefore < 0 {
		return errors.New("padding_digits_before must be a positive integer")
	}
	if gc.PaddingDigitsAfter < 0 {
		return errors.New("padding_digits_after must be a positive integer")
	}

	if gc.PadToLength < 0 {
		return errors.New("pad_to_length must be a positive integer")
	}

	if gc.PaddingCharactersBefore < 0 {
		return errors.New("padding_characters_before must be a positive integer")
	}
	if gc.PaddingCharactersAfter < 0 {
		return errors.New("padding_characters_after must be a positive integer")
	}

	return nil
}
