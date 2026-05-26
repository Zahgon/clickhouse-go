package column

import "time"

// getTimeWithDifferentLocation returns the same time but with different location, e.g.
// "2024-08-15 13:22:34 -03:00" will become "2024-08-15 13:22:34 +04:00".
func getTimeWithDifferentLocation(t time.Time, loc *time.Location) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}
