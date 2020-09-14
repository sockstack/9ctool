package log

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type errorLogger struct {
	logger *log.Logger
	once   sync.Once
}

func (i *errorLogger) logOut(format *string, v ...interface{}) {
	i.once.Do(func() {
		i.init()
	})
	if format != nil {
		i.logger.Output(3, fmt.Sprintf(*format, v...))
		return
	}
	i.logger.Output(3, fmt.Sprintln(v...))
}

func (i *errorLogger) init() {
	i.logger = log.New(os.Stdout, "[ERROR] >> ", log.Lmsgprefix|log.Lshortfile|log.Lmicroseconds|log.Ldate)
}
