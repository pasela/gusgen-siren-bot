package main

import (
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const ENV_CONF = "GUSGEN_SIREN_BOT_CONF"

type Config struct {
	Notifier string        `mapstructure:"notifier"`
	Twitter  TwitterConfig `mapstructure:"twitter"`
}

type TwitterConfig struct {
	ConsumerKey    string `mapstructure:"consumer_key"`
	ConsumerSecret string `mapstructure:"consumer_secret"`
	AccessToken    string `mapstructure:"access_token"`
	AccessSecret   string `mapstructure:"access_secret"`
}

func (c *Config) PFlags() *pflag.FlagSet {
	f := pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)

	f.String("notifier", "stdout", "Specify notifier (stdout,twitter)")

	return f
}

func (c *Config) Viper() *viper.Viper {
	v := viper.New()

	conf := os.Getenv(ENV_CONF)
	if conf != "" {
		v.SetConfigFile(conf)
	} else {
		v.SetConfigName("gusgen-siren-bot")
		v.AddConfigPath("./conf")
		v.AddConfigPath(".")
	}

	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))

	v.SetDefault("notifier", "stdout")
	v.BindEnv("notifier")

	v.BindEnv("twitter.consumer_key")
	v.BindEnv("twitter.consumer_secret")
	v.BindEnv("twitter.access_token")
	v.BindEnv("twitter.access_secret")

	return v
}

func InitConfig() Config {
	cfg := Config{}

	f := cfg.PFlags()
	f.Parse(os.Args[1:])

	v := cfg.Viper()
	v.BindPFlags(f)

	conf := os.Getenv(ENV_CONF)
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok && conf == "" {
			// ignore
		} else {
			panic(err)
		}
	}

	v.Unmarshal(&cfg)
	// v.Unmarshal(&cfg.Twitter)

	return cfg
}
