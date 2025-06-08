package set

import (
	"encoding/json"
)

type Set[V comparable] map[V]struct{}

func MakeSet[V comparable](vs ...V) Set[V] {
	s := make(Set[V])
	for _, v := range vs {
		s.Add(v)
	}
	return s
}

func (s Set[V]) Contains(v V) bool {
	_, ok := s[v]
	return ok
}

func (s Set[V]) Add(v V) {
	s[v] = struct{}{}
}

func (s Set[V]) Remove(v V) {
	delete(s, v)
}

func (s Set[V]) Merge(other Set[V]) {
	for v := range other {
		s.Add(v)
	}
}

func (s Set[V]) IsSubset(subset Set[V]) bool {
	for v := range subset {
		if !s.Contains(v) {
			return false
		}
	}
	return true
}

func (s Set[V]) Equals(other Set[V]) bool {
	return s.IsSubset(other) && other.IsSubset(s)
}

func (s Set[V]) DoesIntersect(other Set[V]) bool {
	for v := range s {
		if other.Contains(v) {
			return true
		}
	}
	return false
}

func (s Set[V]) Intersect(other Set[V]) Set[V] {
	n := MakeSet[V]()
	for v := range s {
		if other.Contains(v) {
			n.Add(v)
		}
	}
	for v := range other {
		if s.Contains(v) {
			n.Add(v)
		}
	}
	return n
}

func (s Set[V]) Union(other Set[V]) Set[V] {
	n := MakeSet[V]()
	for v := range s {
		n.Add(v)
	}
	for v := range other {
		n.Add(v)
	}
	return n
}

func (s Set[V]) Minus(other Set[V]) Set[V] {
	n := MakeSet[V]()
	for v := range s {
		if !other.Contains(v) {
			n.Add(v)
		}
	}
	return n
}

func (s Set[V]) Slice() []V {
	l := make([]V, 0, len(s))
	for v := range s {
		l = append(l, v)
	}
	return l
}

func (s Set[V]) UnmarshalJSON(bytes []byte) error {
	var list []V
	err := json.Unmarshal(bytes, &list)
	if err != nil {
		return err
	}
	for _, v := range list {
		s.Add(v)
	}
	return nil
}

func (s Set[V]) MarshalJSON() ([]byte, error) {
	list := make([]V, 0, len(s))
	for v := range s {
		list = append(list, v)
	}
	bytes, err := json.Marshal(&list)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
