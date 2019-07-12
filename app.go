package main

import (
	"errors"
	"log"

	"github.com/pasela/go-vanatime"
	"github.com/pasela/gusgen-siren-bot/notifier"
	"github.com/pasela/gusgen-siren-bot/timer"
)

type App struct {
	ticker   *timer.Ticker
	notifier notifier.Notifier
	logger   *log.Logger
	done     chan struct{}
}

func NewApp(notifier notifier.Notifier, logger *log.Logger) *App {
	app := &App{
		ticker:   timer.NewTicker(),
		notifier: notifier,
		logger:   logger,
		done:     make(chan struct{}),
	}

	return app
}

func (a *App) Run() error {
	a.logger.Printf("next: %s / %s", a.ticker.Next, a.ticker.Next.Earth())
	for {
		select {
		case value, ok := <-a.ticker.C:
			if ok {
				go a.notify(value)
				a.logger.Printf("next: %s / %s", a.ticker.Next, a.ticker.Next.Earth())
			} else {
				return errors.New("Ticker closed")
			}

		case <-a.done:
			return nil
		}
	}
}

func (a *App) notify(vt vanatime.Time) {
	if err := a.notifier.Notify(vt); err != nil {
		a.logger.Println("Notification error:", err)
	}
}

func (a *App) Stop() {
	close(a.done)
}
