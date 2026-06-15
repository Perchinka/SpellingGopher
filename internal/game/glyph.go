package game

type CharState int

const (
	Pending CharState = iota
	Correct
	Wrong
)

type Glyph struct {
	Expected rune
	Current  rune
	State    CharState
}

func (g *Glyph) IsSpace() bool {
	return g.Current == ' '
}
