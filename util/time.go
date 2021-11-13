package util

import "time"

func GetTimeNow() time.Time {
	return time.Now().In(time.UTC)
}
