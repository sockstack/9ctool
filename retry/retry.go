package retry

import "time"

//Retry 执行可重试任务
func Retry(callback func() error, max int, interval time.Duration) (err error) {
	for i := 0; i < max; i++ {
		if err = callback(); err != nil {
			time.Sleep(interval)
			continue
		}

		return
	}

	return
}
