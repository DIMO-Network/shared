package set

// Set is a set of elements.
type Set[A comparable] map[A]struct{}

// New creates a new set.
func New[A comparable](as ...A) Set[A] {
	return make(Set[A])
}

// Add adds an element to the set.
func (s Set[A]) Add(a A) {
	s[a] = struct{}{}
}

// Contains checks if an element is in the set.
func (s Set[A]) Contains(a A) bool {
	_, ok := s[a]
	return ok
}
