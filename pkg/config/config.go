package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	// DefaultAlphabet is the default character set used for the separator
	// and padding modifications.
	DefaultAlphabet = []string{
		"!", "@", "$", "%", "^", "&", "*", "-", "_",
		"+", "=", ":", "|", "~", "?", "/", ".", ";",
	}

	// ErrNumWordsNegative is returned when num_words is negative.
	ErrNumWordsNegative = errors.New("num_words must be a positive integer")

	// ErrWordLenMinNegative is returned when word_len_min is negative.
	ErrWordLenMinNegative = errors.New("word_len_min must be a positive integer")

	// ErrWordLenMaxNegative is returned when word_len_max is negative.
	ErrWordLenMaxNegative = errors.New("word_len_max must be a positive integer")

	// ErrWordLenMaxLessThanMin is returned when word_len_max is less than
	// word_len_min.
	ErrWordLenMaxLessThanMin = errors.New("word_len_max cannot be less than word_len_min")

	// ErrPaddingDigitsBeforeNeg is returned when padding_digits_before is
	// negative.
	ErrPaddingDigitsBeforeNeg = errors.New("padding_digits_before must be a positive integer")

	// ErrPaddingDigitsAfterNeg is returned when padding_digits_after is
	// negative.
	ErrPaddingDigitsAfterNeg = errors.New("padding_digits_after must be a positive integer")

	// ErrPadToLengthNegative is returned when pad_to_length is negative.
	ErrPadToLengthNegative = errors.New("pad_to_length must be a positive integer")

	// ErrPaddingCharsBeforeNeg is returned when padding_characters_before is
	// negative.
	ErrPaddingCharsBeforeNeg = errors.New("padding_characters_before must be a positive integer")

	// ErrPaddingCharsAfterNeg is returned when padding_characters_after is
	// negative.
	ErrPaddingCharsAfterNeg = errors.New("padding_characters_after must be a positive integer")
)

// GeneratorConfig represents the configuration for the password generator.
type GeneratorConfig struct {
	NumWords                int      `mapstructure:"num_words"`
	WordLenMin              int      `mapstructure:"word_length_min"`
	WordLenMax              int      `mapstructure:"word_length_max"`
	CaseTransform           string   `mapstructure:"case_transform"`
	SeparatorCharacter      string   `mapstructure:"separator_character"`
	SeparatorAlphabet       []string `mapstructure:"separator_alphabet"`
	PaddingDigitsBefore     int      `mapstructure:"padding_digits_before"`
	PaddingDigitsAfter      int      `mapstructure:"padding_digits_after"`
	PaddingType             string   `mapstructure:"padding_type"`
	PaddingCharacter        string   `mapstructure:"padding_character"`
	SymbolAlphabet          []string `mapstructure:"symbol_alphabet"`
	PadToLength             int      `mapstructure:"pad_to_length"`
	PaddingCharactersBefore int      `mapstructure:"padding_characters_before"`
	PaddingCharactersAfter  int      `mapstructure:"padding_characters_after"`
}

func init() {
	pflag.StringVar(&cfgFile, "config", mustDefaultConfigFile(), "path to config file")

	// Define flags with names matching the mapstructure tags on the config
	// struct. This allows Viper to override config file values with those
	// from the flags.
	pflag.Int("num_words", 3, "number of words")
	pflag.Int("word_length_min", 4, "minimum word length")
	pflag.Int("word_length_max", 8, "maximum word length")
	pflag.String(
		"case_transform", "ALTERNATE",
		"case transformation, allowed values: LOWER, UPPER, RANDOM, NONE, ALTERNATE, CAPITALISE, INVERT",
	)
	pflag.String("separator_character", "RANDOM", "character to separate password parts")
	pflag.StringSlice(
		"separator_alphabet",
		DefaultAlphabet,
		"comma-separated list of characters to separate password parts",
	)
	pflag.Int("padding_digits_before", 2, "number of digits to pad before the password")
	pflag.Int("padding_digits_after", 2, "number of digits to pad before the password")
	pflag.String(
		"padding_type", "FIXED", "padding type, allowed values: FIXED, ADAPTIVE",
	)
	pflag.String("padding_character", "RANDOM", "character to pad the password with")
	pflag.StringSlice(
		"symbol_alphabet",
		DefaultAlphabet,
		"comma-separated list of characters to pad the password with",
	)
	pflag.Int("pad_to_length", 8, "length to pad the password to, will be ignored if less than the generated password length")
	pflag.Int("padding_characters_before", 2, "number of characters to pad before the password")
	pflag.Int("padding_characters_after", 2, "number of characters to pad before the password")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)
}

// mustDefaultConfigFile returns the path to the default config file.
//
// The default config file is expected to exist in the users home directory.
//
// If an error is encountered finding the home directory, the method will panic.
func mustDefaultConfigFile() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return home + "/.xkpassgo.json"
}

// Load returns a fully loaded configuration for the password generator.
//
// The config will be taken from the defined config file (default or from flag)
// and can be overridden with flag values.
func Load() (*GeneratorConfig, error) {
	viper.SetConfigType("json")

	// Only load config from file if it exists.
	_, err := os.Stat(cfgFile)
	switch {
	case err == nil:
		viper.SetConfigFile(cfgFile)

		if err := viper.ReadInConfig(); err != nil {
			return nil, fmt.Errorf("read config: %w", err)
		}
	case os.IsNotExist(err):
		// Config file does not exist. Do nothing.
	default:
		return nil, fmt.Errorf("config file exists: %w", err)
	}

	var cfg GeneratorConfig

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	return &cfg, nil
}

// Validate returns an error if the current configuration is not valid.
func (gc *GeneratorConfig) Validate() error {
	if gc.NumWords < 0 {
		return ErrNumWordsNegative
	}

	if gc.WordLenMin < 0 {
		return ErrWordLenMinNegative
	}

	if gc.WordLenMax < 0 {
		return ErrWordLenMaxNegative
	}

	if gc.WordLenMax < gc.WordLenMin {
		return ErrWordLenMaxLessThanMin
	}

	if gc.PaddingDigitsBefore < 0 {
		return ErrPaddingDigitsBeforeNeg
	}

	if gc.PaddingDigitsAfter < 0 {
		return ErrPaddingDigitsAfterNeg
	}

	if gc.PadToLength < 0 {
		return ErrPadToLengthNegative
	}

	if gc.PaddingCharactersBefore < 0 {
		return ErrPaddingCharsBeforeNeg
	}

	if gc.PaddingCharactersAfter < 0 {
		return ErrPaddingCharsAfterNeg
	}

	return nil
}
