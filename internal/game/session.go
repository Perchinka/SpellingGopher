package game

import "time"

// NOTE not sure if this CharState has to be here
// I guess for now it's fine, but probably will move later

type CharState int

const (
	Pending CharState = iota
	Correct
	Wrong
)

type Glyph struct {
	expected rune
	current  rune
	state    CharState
}

type Session struct {
	target  []rune
	typed   []rune
	started time.Time
	ended   time.Time
}

func NewSession(target string) Session {
	return Session{
		target:  []rune(target),
		started: time.Now(),
	}
}

func (s *Session) Type(r rune) {
	if len(s.typed) >= len(s.target) {
		return
	}
	s.typed = append(s.typed, r)
}

func (s *Session) glyphAt(i int, expected rune) Glyph {
	if i >= len(s.typed) {
		return Glyph{expected: expected, current: expected, state: Pending}
	}
	typed := s.typed[i]
	if typed == expected {
		return Glyph{expected: expected, current: typed, state: Correct}
	}
	return Glyph{expected: expected, current: typed, state: Wrong}
}

func (s *Session) Glyphs() []Glyph {
	glyphs := make([]Glyph, len(s.target))
	for i, expected := range s.target {
		glyphs[i] = s.glyphAt(i, expected)
	}
	return glyphs
}
