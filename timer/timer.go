package timer

import (
	"sync"

	"github.com/pasela/go-vanatime"
)

type Ticker struct {
	C    <-chan vanatime.Time
	Next vanatime.Time

	c     chan<- vanatime.Time
	timer *vanatime.Timer
	stop  chan struct{}
	wg    sync.WaitGroup
}

func NewTicker() *Ticker {
	next := nextTime()
	timer := vanatime.NewTimer(next.Sub(vanatime.Now()))
	c := make(chan vanatime.Time)

	ticker := &Ticker{
		C:     c,
		Next:  next,
		c:     c,
		timer: timer,
		stop:  make(chan struct{}),
	}

	ticker.start()

	return ticker
}

func (t *Ticker) start() {
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		for {
			select {
			case value, ok := <-t.timer.C:
				if ok {
					t.c <- value
					t.update()
				} else {
					close(t.c)
					return
				}

			case <-t.stop:
				return
			}
		}
	}()
}

func nextTime() vanatime.Time {
	return vanatime.Now().AddDate(0, 0, 1).Truncate(vanatime.Day)
}

func (t *Ticker) update() {
	next := nextTime()
	d := next.Sub(vanatime.Now())
	t.Next = next
	t.timer.Reset(d)
}

func (t *Ticker) Stop() {
	if t.timer != nil {
		t.timer.Stop()
		if t.stop != nil {
			close(t.stop)
			t.wg.Wait()
		}
	}
}
