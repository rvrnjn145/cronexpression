package cronExpression

import "fmt"

var (
	key = map[int]string{
		0: "minute",
		1: "hour",
		2: "day of month",
		3: "month",
		4: "day of week",
		5: "command",
	}
)

// Logger is the logging interface for expression descriptor.
type Logger interface {
	Printf(format string, v []string)
}

type loggerImpl struct {
	LoggerFlag bool
}

func (l *loggerImpl) Printf(formmat string, v []string) {
	if l.LoggerFlag {
		for i, val := range v {
			fmt.Printf("%-14s%s\n", key[i], val)
		}
	}
}
