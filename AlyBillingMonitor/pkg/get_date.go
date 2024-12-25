package pkg

import (
	"strings"
	"time"
)

// getFirstDateOfMonth 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func getFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return getZeroTime(d)
}

// getLastDateOfMonth 获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func getLastDateOfMonth(d time.Time) time.Time {
	return getFirstDateOfMonth(d).AddDate(0, 1, -1)
}

// getZeroTime 获取某一天的0点时间
func getZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// GetAZDay 获取
func GetAZDay(d time.Time) (string, string) {
	firstDate := strings.Split(getFirstDateOfMonth(d).String(), " ")[0]
	LastDate := strings.Split(getLastDateOfMonth(d).String(), " ")[0]

	// 形如：2022-09-01
	return firstDate, LastDate
}
