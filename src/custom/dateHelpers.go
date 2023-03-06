package custom

import "time"

const (
	UtcLocation                   = "UTC"
	DateLayout                    = "2006-01-02"
	DateTimeLayout                = "2006-01-02 15:04:05"
	DateTimeWithNumericZone       = "2006-01-02 15:04:05 -0700"
	DateTimeWithNumericAndUtcZone = "2006-01-02 15:04:05 -0700 MST"
	DateDisplay                   = "Monday, 02 Jan 2006"
	DateTimeDisplay               = "Monday, 02 Jan 2006 15:04:05"
)

func setLocalTimeFromTime(t time.Time) time.Time {
	return t.Add(calculateTimeDifferenceUTC())
}

func calculateTimeDifferenceUTC() time.Duration {
	_, offset := time.Now().Zone()
	return time.Duration(offset) * time.Second
}
