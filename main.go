package main

import (
	"log"
	"os"
	"strings"
	"sync"
	"syscall"

	"github.com/pasela/go-vanatime"
	"github.com/pasela/gusgen-siren-bot/notifier"
	"github.com/pasela/gusgen-siren-bot/sig"
)

func main() {
	cfg := InitConfig()
	run(cfg)
}

func initLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags)
}

func initNotifier(cfg Config, logger *log.Logger) notifier.Notifier {
	var n notifier.Notifier

	switch strings.ToLower(cfg.Notifier) {
	case "stdout":
		n = notifier.NewStdoutNotifier()
		n.SetFormatter(new(notifier.SimpleFormatter))

	case "twitter":
		twCfg := notifier.TwitterNotifierConfig{
			ConsumerKey:    cfg.Twitter.ConsumerKey,
			ConsumerSecret: cfg.Twitter.ConsumerSecret,
			AccessToken:    cfg.Twitter.AccessToken,
			AccessSecret:   cfg.Twitter.AccessSecret,
		}
		var err error
		n, err = notifier.NewTwitterNotifier(twCfg, logger)
		if err != nil {
			panic("Failed to initialize Twitter notifier: " + err.Error())
		}
		n.SetFormatter(new(notifier.TweetFormatter))

	default:
		panic("Unsupported notifier: " + cfg.Notifier)
	}

	return n
}

func run(cfg Config) {
	logger := initLogger()
	now := vanatime.Now()
	logger.Printf("Startup at %s / %s", now, now.Earth())
	defer logger.Println("Stop")

	notifier := initNotifier(cfg, logger)
	app := NewApp(vanatime.Day, notifier, logger)

	var once sync.Once
	cancel := sig.Handle(func(sig os.Signal) {
		logger.Println("Signal received:", sig)
		once.Do(func() {
			app.Stop()
		})
	}, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := app.Run(); err != nil {
		logger.Println(err)
	}
}
