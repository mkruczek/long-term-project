package trade

import "time"

type Filter struct {
	StartTime time.Time
	EndTime   time.Time
	Symbol    string
}
