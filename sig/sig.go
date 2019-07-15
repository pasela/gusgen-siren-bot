package sig

import (
	"context"
	"os"
	"os/signal"
)

type SignalFunc func(os.Signal)

func Handle(callback SignalFunc, signals ...os.Signal) context.CancelFunc {
	return HandleContext(context.Background(), callback, signals...)
}

func HandleContext(ctx context.Context, callback SignalFunc, signals ...os.Signal) context.CancelFunc {
	c, cancel := context.WithCancel(ctx)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, signals...)

	go func() {
		for {
			select {
			case sig := <-sigCh:
				callback(sig)
			case <-c.Done():
				return
			}
		}
	}()

	return cancel
}
