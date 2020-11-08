package cor

import (
	"fmt"
)

type Resolver interface {
	resolve(*Trouble) bool
}

type SupportMember struct {
	Name string
	next *SupportMember
	Resolver
}

func (s *SupportMember) SetNext(next *SupportMember) *SupportMember {
	s.next = next
	return s
}

func (s *SupportMember) Support(t *Trouble) {
	fmt.Println(s.Name, t.string())
	if s.resolve(t) {
		s.done(t)
	} else if s.next != nil {
		s.next.Support(t)
	} else {
		s.fail(t)
	}
}

func (s *SupportMember) string() string {
	return fmt.Sprintf("[%v]", s.Name)
}

func (s *SupportMember) done(t *Trouble) {
	fmt.Printf("%v is resolved by %v.", t.string(), s.string())
}

func (s *SupportMember) fail(t *Trouble) {
	fmt.Printf("%v cannot be resolved.", t.string())
}
