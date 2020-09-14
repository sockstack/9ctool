package log

import (
	"github.com/sockstack/9ctool/log"
	"testing"
)

func TestInfo(t *testing.T) {
	log.Info("ok")
}

func TestFInfo(t *testing.T)  {
	log.FInfo("%s,%d", "ok", 123)
}
