package notifier

import (
	"errors"
	"io/ioutil"
	"log"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/pasela/go-vanatime"
)

var _ Notifier = (*TwitterNotifier)(nil)

type TwitterNotifierConfig struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

type TwitterNotifier struct {
	api       *anaconda.TwitterApi
	formatter Formatter
	logger    *log.Logger
}

func (n *TwitterNotifier) SetFormatter(formatter Formatter) {
	n.formatter = formatter
}

func NewTwitterNotifier(cfg TwitterNotifierConfig, logger *log.Logger) (Notifier, error) {
	api := anaconda.NewTwitterApiWithCredentials(
		cfg.AccessToken,
		cfg.AccessSecret,
		cfg.ConsumerKey,
		cfg.ConsumerSecret,
	)

	if logger == nil {
		logger = log.New(ioutil.Discard, "", 0)
	}

	user, err := api.GetSelf(nil)
	if err != nil {
		return nil, err
	}
	logger.Printf("Verify credentials ...ok (id=%s, screen_name=%s)",
		user.IdStr, user.ScreenName)

	return &TwitterNotifier{
		api:    api,
		logger: logger,
	}, nil
}

func (n *TwitterNotifier) Notify(vt vanatime.Time) error {
	limit := vt.Add(vanatime.Hour * 12)
	maxRetry := 5
	sleep := time.Minute * 3

	return retryUntil(func(t vanatime.Time) error {
		_, err := n.api.PostTweet(format(n.formatter, t), nil)
		return err
	}, vt, limit, maxRetry, sleep, n.logger)
}

func retryUntil(block func(vanatime.Time) error, initial vanatime.Time, limit vanatime.Time, maxRetry int, sleep time.Duration, logger *log.Logger) error {
	var vt = initial
	var count int
	for {
		count++
		if err := block(vt); err == nil {
			return nil
		} else {
			logger.Println(err)
		}

		if vanatime.Now().Before(limit) && count <= maxRetry {
			logger.Printf("Retrying after %s (%d/%d)", sleep, count, maxRetry)
			time.Sleep(sleep)
		} else {
			return errors.New("Retry over")
		}
		vt = vanatime.Now()
	}
}
