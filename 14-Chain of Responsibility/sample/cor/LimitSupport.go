package cor

type LimitSupport struct {
	Limit int
}

func (s *LimitSupport) resolve(t *Trouble) bool {
	return t.Number < s.Limit
}
