package quote

import "unicode/utf8"

type Quote struct {
	Text   string
	Author string
}

func (q Quote) RuneCount() int {
	return utf8.RuneCountInString(q.Text)
}
