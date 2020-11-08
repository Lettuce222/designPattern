package cor

type specialSupport struct {
	SupportMember
	number int
}

func NewSpecialSupport(Name string, number int) *specialSupport {
	return &specialSupport{SupportMember: SupportMember{Name: Name}, number: number}
}

func (s *specialSupport) resolve(t *Trouble) bool {
	return t.Number == s.number
}
