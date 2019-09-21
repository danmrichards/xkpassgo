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

type GeneratorConfig struct {
	NumWords      int             `mapstructure:"num_words"`
	WordLenMin    int             `mapstructure:"word_length_min"`
	WordLenMax    int             `mapstructure:"word_length_max"`
	CaseTransform transform.Style `mapstructure:"case_transform"`
}

func init() {
	pflag.StringVar(&cfgFile, "config", mustDefaultConfigFile(), "path to config file")
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

func mustDefaultConfigFile() string {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	return home + "/.xkpassgo.json"
}

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
