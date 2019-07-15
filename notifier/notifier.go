package notifier

import "github.com/pasela/go-vanatime"

type Notifier interface {
	Notify(vanatime.Time) error
	SetFormatter(formatter Formatter)
}

func format(formatter Formatter, vt vanatime.Time) string {
	if formatter == nil {
		return vt.String()
	}
	return formatter.Format(vt)
}
