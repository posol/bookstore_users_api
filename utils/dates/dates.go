package dates

import (
	"time"
)

const (
	apiDateLayout = "Mon Jan 2 15:04:05 -0700 MST 2006"
	apiDbLayout   = "2006-01--2 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
func GetNowDBFormat() string {
	return GetNow().Format(apiDbLayout)
}
