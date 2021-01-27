package entity

import "time"

const Day = time.Hour * 24

// Timestamp returns the current time in UTC rounded to seconds.
func Timestamp() time.Time {
	return time.Now().UTC().Round(time.Second)
}

// Seconds converts an int to a duration in seconds.
func Seconds(s int) time.Duration {
	return time.Duration(s) * time.Second
}

// Yesterday returns the time 24 hours ago.
func Yesterday() time.Time {
	return time.Now().Add(-24 * time.Hour)
}
