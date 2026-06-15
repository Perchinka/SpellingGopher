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

func TestSession_Accuracy(t *testing.T) {
	const tol = 1e-9
	assertAccuracy := func(t *testing.T, s *Session, want float64) {
		t.Helper()
		if got := s.Accuracy(); math.Abs(got-want) > tol {
			t.Errorf("Accuracy() = %v, want %v", got, want)
		}
	}

	t.Run("fresh session is 100% (guards 0/0)", func(t *testing.T) {
		s := NewSession("cat")
		assertAccuracy(t, s, 1.0)
	})

	t.Run("all correct is 100%", func(t *testing.T) {
		s := NewSession("cat")
		typeString(s, "cat")
		assertAccuracy(t, s, 1.0)
	})

	t.Run("one wrong of three is two thirds", func(t *testing.T) {
		s := NewSession("cat")
		typeString(s, "cax")
		assertAccuracy(t, s, 2.0/3.0)
	})

	t.Run("all wrong is 0%", func(t *testing.T) {
		s := NewSession("cat")
		typeString(s, "xyz")
		assertAccuracy(t, s, 0.0)
	})

	t.Run("backspace does not forgive a past mistake", func(t *testing.T) {
		// The cumulative model: type a wrong char, backspace it, type the right
		// one. The final text "ca" is clean, but the mistake still happened.
		// 3 keystrokes, 1 error -> 2/3. Backspace edits the text, not history.
		s := NewSession("cat")
		typeString(s, "cx")
		s.Backspace()
		s.Type('a')
		assertAccuracy(t, s, 2.0/3.0)
	})

	t.Run("overtyping is dropped and does not count", func(t *testing.T) {
		s := NewSession("cat")
		typeString(s, "catxy")
		assertAccuracy(t, s, 1.0)
	})
	t.Run("retyping a correct char keeps 100%", func(t *testing.T) {
		s := NewSession("cat")
		s.Type('c')
		s.Backspace()
		s.Type('c')
		assertAccuracy(t, s, 1.0)
	})
}
