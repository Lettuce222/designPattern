package cor

type oddSupport struct {
	SupportMember
}

func NewOddSupport(Name string) *oddSupport {
	return &oddSupport{SupportMember: SupportMember{Name: Name}}
}

func (s *oddSupport) resolve(t *Trouble) bool {
	return t.Number%2 == 1
}
