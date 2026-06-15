package game

import (
	"math"
	"testing"
	"time"
)

func TestWPM(t *testing.T) {
	tests := []struct {
		name       string
		characters int
		duration   time.Duration
		want       float64
	}{
		{
			name:       "100 wpm",
			characters: 500,
			duration:   time.Minute,
			want:       100,
		},
		{
			name:       "sub-minute scales up",
			characters: 25,
			duration:   30 * time.Second,
			want:       10,
		},
		{
			name:       "no characters is zero",
			characters: 0,
			duration:   time.Minute,
			want:       0,
		},
		{
			name:       "zero duration is guarded, not NaN",
			characters: 100,
			duration:   0,
			want:       0,
		},
		{
			name:       "negative duration is guarded",
			characters: 100,
			duration:   -time.Second,
			want:       0,
		},
		{
			name:       "zero over zero is not NaN",
			characters: 0,
			duration:   0,
			want:       0,
		},
		{
			name:       "smallest positive duration stays finite",
			characters: 1,
			duration:   time.Nanosecond,
			want:       (1.0 / 5.0) / time.Nanosecond.Minutes(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WPM(tt.characters, tt.duration)
			if math.IsNaN(got) || math.IsInf(got, 0) {
				t.Fatalf("WPM(%d, %v) = %v, want a finite number", tt.characters, tt.duration, got)
			}
			if got != tt.want {
				t.Errorf("WPM(%d, %v) = %v, want %v", tt.characters, tt.duration, got, tt.want)
			}
		})
	}
}
