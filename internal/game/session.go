package game

import "time"

type Session struct {
	target  []rune
	typed   []rune
	started time.Time
	ended   time.Time
	clock   Clock
}

func NewSession(target string) *Session {
	return &Session{
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

func (s *Session) Backspace() {
	if len(s.typed) <= 0 {
		return
	}
	s.typed = s.typed[:len(s.typed)-1]
}

func (s *Session) glyphAt(i int, expected rune) Glyph {
	if i >= len(s.typed) {
		return Glyph{Expected: expected, Current: expected, State: Pending}
	}
	typed := s.typed[i]
	if typed == expected {
		return Glyph{Expected: expected, Current: typed, State: Correct}
	}
	return Glyph{Expected: expected, Current: typed, State: Wrong}
}

func (s *Session) Glyphs() []Glyph {
	glyphs := make([]Glyph, len(s.target))
	for i, expected := range s.target {
		glyphs[i] = s.glyphAt(i, expected)
	}
	return glyphs
}

func (s *Session) Finished() bool {
	return len(s.typed) == len(s.target)
}
