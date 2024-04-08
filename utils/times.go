package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type HTime struct {
	time.Time
}

var formatTime = "2006-01-02 13:04:05"

func (t HTime) MarshalJSON() ([]byte, error) {
	formatted := t.Format(formatTime)
	return []byte(formatted), nil
}

func (t *HTime) UnmarshalJSON(data []byte) (err error) {
	now, _ := time.Parse(formatTime, string(data))
	*t = HTime{Time: now}
	return err
}
func (t HTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if zeroTime.UnixNano() == t.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *HTime) Scan(v any) error {
	value, ok := v.(time.Time)
	if ok {
		*t = HTime{Time: value}
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
