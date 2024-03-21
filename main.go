package main

type F interface {
	f()
}
type s1 struct{}

func (s s1) f() {}

type s2 struct {
}

func (s *s2) f() {
}

func msin() {
	s1Val := s1{}
	s1Ptr := &s1{}
	s21 := s2{}
	s22 := &s2{}

	var i F
	i = s1Val
	i = s1Ptr
	i = s21
	i = s22

	_ = i

}
