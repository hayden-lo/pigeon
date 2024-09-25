package utils

import (
	"time"
)

func GetNowTime() string {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return "Error loading location"
	}
	now := time.Now().In(loc)
	return now.Format("2006-01-02 15:04:05")
}

func StringToTime(s string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, s)
	return t, err
}
