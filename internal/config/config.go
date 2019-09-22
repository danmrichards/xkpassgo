package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/danmrichards/xkpassgo/internal/transform"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var cfgFile string

// GeneratorConfig represents the configuration for the password generator.
type GeneratorConfig struct {
	NumWords      int             `mapstructure:"num_words"`
	WordLenMin    int             `mapstructure:"word_length_min"`
	WordLenMax    int             `mapstructure:"word_length_max"`
	CaseTransform transform.Style `mapstructure:"case_transform"`
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
		"case_transform",
		"",
		"case transformation, allowed values: "+strings.Join(transform.Styles(), ", "),
	)
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)
}

// mustDefaultConfigFile returns the path to the default config file.
//
// The default config file is expected to exist in the users home directory.
//
// If an error is encountered finding the home directory, the method will panic.
func mustDefaultConfigFile() string {
	home, err := homedir.Dir()
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
	switch _, err := os.Stat(cfgFile); {
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
