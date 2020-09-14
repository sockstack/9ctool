package log

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type debugLogger struct {
	logger *log.Logger
	once   sync.Once
}

func (i *debugLogger) logOut(format *string, v ...interface{}) {
	i.once.Do(func() {
		i.init()
	})
	if format != nil {
		i.logger.Output(3, fmt.Sprintf(*format, v...))
		return
	}
	i.logger.Output(3, fmt.Sprintln(v...))
}

func (i *debugLogger) init() {
	i.logger = log.New(os.Stdout, "[DEBUG] >> ", log.Lmsgprefix|log.Lshortfile|log.Lmicroseconds|log.Ldate)
}
