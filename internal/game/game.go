package game

import "time"

func WPM(characters int, duration time.Duration) float64 {
	if duration <= 0 {
		return 0.0
	}

	return (float64(characters) / 5.0) / duration.Minutes()
}
