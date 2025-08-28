package utils

// implements set with map
type Set[T comparable] struct {
	data map[T]bool
}

func New[T comparable]() *Set[T] {
	return &Set[T]{data: make(map[T]bool)}
}

func (s *Set[T]) Clear() {
	s.data = make(map[T]bool)
}

func (s *Set[T]) Add(val T) {
	s.data[val] = true
}

func (s *Set[T]) Remove(val T) {
	delete(s.data, val)
}

func (s Set[T]) Contains(val T) bool {
	_, ok := s.data[val]
	return ok
}

func (s Set[T]) Len() int {
	return len(s.data)
}

func (s Set[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s Set[T]) IsNotEmpty() bool {
	return !s.IsEmpty()
}

func (s Set[T]) Values() []T {
	result := make([]T, 0, len(s.data))
	for k := range s.data {
		result = append(result, k)
	}
	return result
}

func (s Set[T]) Union(other *Set[T]) *Set[T] {
	res := New[T]()
	for k := range s.data {
		res.Add(k)
	}
	for k := range other.data {
		res.Add(k)
	}
	return res
}

func (s Set[T]) Intersect(other *Set[T]) *Set[T] {
	res := New[T]()
	for k := range s.data {
		if other.Contains(k) {
			res.Add(k)
		}
	}
	return res
}

func (s Set[T]) Difference(other *Set[T]) *Set[T] {
	res := New[T]()
	for k := range s.data {
		if !other.Contains(k) {
			res.Add(k)
		}
	}
	return res
}
