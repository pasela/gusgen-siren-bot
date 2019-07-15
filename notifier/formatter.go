package notifier

import (
	"fmt"

	"github.com/pasela/go-vanatime"
)

type Formatter interface {
	Format(vanatime.Time) string
}

type SimpleFormatter struct {
}

func (f *SimpleFormatter) Format(vt vanatime.Time) string {
	return vt.String()
}

type TweetFormatter struct {
}

func (f *TweetFormatter) Format(vt vanatime.Time) string {
	return "ウゥーーーーーーーーーーーーーーー\n" +
		formatVanaTime(vt)
}

func formatVanaTime(vt vanatime.Time) string {
	return fmt.Sprintf(
		"天晶暦%s %s %s",
		vt.Strftime("%Y/%m/%d %H:%M:%S"),
		vt.Weekday().StringLocale("ja"),
		vt.Moon().Phase().StringLocale("ja"),
	)
}
