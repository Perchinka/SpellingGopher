package game

import (
	"testing"
	"time"
)

func TestWPM(t *testing.T) {
	tests := []struct {
		name    string
		chars   int
		elapsed time.Duration
		want    float64
	}{
		{name: "60 chars in 60s is 12wpm", chars: 60, elapsed: time.Minute, want: 12},
		{name: "zero time", chars: 60, elapsed: 0, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WPM(tt.chars, tt.elapsed)
			if got != tt.want {
				t.Errorf("WPM(%d, %v) = %v, want %v", tt.chars, tt.elapsed, got, tt.want)
			}
		})
	}
}
