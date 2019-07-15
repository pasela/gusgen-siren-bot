package notifier

import (
	"fmt"
	"io"
	"os"

	"github.com/pasela/go-vanatime"
)

var _ Notifier = (*WriterNotifier)(nil)

type WriterNotifier struct {
	w         io.Writer
	formatter Formatter
}

func (n *WriterNotifier) SetFormatter(formatter Formatter) {
	n.formatter = formatter
}

func (n *WriterNotifier) Notify(vt vanatime.Time) error {
	_, err := fmt.Fprintln(n.w, format(n.formatter, vt))
	return err
}

func NewStdoutNotifier() Notifier {
	return &WriterNotifier{
		w: os.Stdout,
	}
}
