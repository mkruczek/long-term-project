package statistics_v2

import "time"

type Filter struct {
	StartTime time.Time
	EndTime   time.Time
	Symbol    string
}
