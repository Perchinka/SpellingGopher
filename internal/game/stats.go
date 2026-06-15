package game

import "time"

type Clock interface{ Now() time.Time }

type RealClock struct{}

func (RealClock) Now() time.Time { return time.Now() }

type Stats struct {
	Elapsed  time.Duration
	WPM      float64
	Accuracy float64
}

func (s *Session) Elapsed() time.Duration {
	if s.started.IsZero() {
		return 0
	}
	if !s.ended.IsZero() {
		return s.ended.Sub(s.started)
	}
	return s.clock.Now().Sub(s.started)
}

func (s *Session) Accuracy() float64 {
	if s.keystrokes == 0 {
		return 1.0
	}
	return float64(s.keystrokes-s.errors) / float64(s.keystrokes)
}

func (s *Session) Stats() Stats {
	return Stats{
		Elapsed:  s.Elapsed(),
		WPM:      WPM(len(s.target), s.Elapsed()), // NOTE Elapsed could be 0
		Accuracy: s.Accuracy(),                    //TODO calculate actual accuracy
	}
}

func WPM(characters int, duration time.Duration) float64 {
	if duration <= 0 {
		return 0.0
	}

	return (float64(characters) / 5.0) / duration.Minutes()
}
