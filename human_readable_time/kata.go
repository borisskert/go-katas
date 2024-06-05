package kata

import (
	"fmt"
	"time"
)

// HumanReadableTime / Human Readable Time
// https://www.codewars.com/kata/52685f7382004e774f0001f7/train/go
func HumanReadableTime(seconds int) string {
	duration := time.Second * time.Duration(seconds)

	hour := int(duration.Hours())
	minute := int(duration.Minutes()) % 60
	second := int(duration.Seconds()) % 60

	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}
