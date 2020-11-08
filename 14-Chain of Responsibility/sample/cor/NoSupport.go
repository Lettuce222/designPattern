package cor

type NoSupport struct {
}

func (s *NoSupport) resolve(t *Trouble) bool {
	return false
}
