package retry

import (
	"errors"
	"github.com/sockstack/9ctool/retry"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	err := errors.New("test retry error")
	i := 1
	assert.Error(t, err, retry.Retry(func() error {
		log.Printf("time: %d", i)
		i++
		return err
	}, 3, time.Second * 1))
}
