package game

import "time"

var fakeEpoch = time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)

type fakeClock struct{ elapsed time.Duration }

func (c *fakeClock) Now() time.Time          { return fakeEpoch.Add(c.elapsed) }
func (c *fakeClock) Advance(d time.Duration) { c.elapsed += d }

func newTestSession(target string) *Session {
	return NewSession(target, &fakeClock{}, "FakeAuthor")
}
