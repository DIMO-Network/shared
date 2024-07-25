package set

type Set[A comparable] map[A]struct{}

func New[A comparable]() Set[A] {
	return make(Set[A])
}

func (s Set[A]) Add(a A) {
	s[a] = struct{}{}
}

func (s Set[A]) Contains(a A) bool {
	_, ok := s[a]
	return ok
}
