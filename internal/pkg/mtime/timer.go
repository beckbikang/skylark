package mtime

import (
	"time"
)

func GetNowTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}
