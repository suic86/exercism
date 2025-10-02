package stringset

import (
	"fmt"
	"strings"
)

type Set struct {
	data map[string]bool
}

func New() Set {
	return Set{
		data: make(map[string]bool),
	}
}

func NewFromSlice(l []string) Set {
	s := New()
	for _, k := range l {
		s.data[k] = true
	}
	return s
}

func (s Set) String() string {
	if len(s.data) == 0 {
		return "{}"
	}
	l := []string{}
	for k := range s.data {
		l = append(l, "\""+k+"\"")
	}
	return fmt.Sprintf("{%v}", strings.Join(l, ", "))
}

func (s Set) IsEmpty() bool {
	return len(s.data) == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s.data[elem]
	return ok
}

func (s Set) Add(elem string) {
	s.data[elem] = true
}

func Subset(s1, s2 Set) bool {
	for k := range s1.data {
		if _, ok := s2.data[k]; !ok {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for k := range s1.data {
		if _, ok := s2.data[k]; ok {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	return Subset(s1, s2) && Subset(s2, s1)
}

func Intersection(s1, s2 Set) Set {
	r := New()
	for k := range s1.data {
		if _, ok := s2.data[k]; ok {
			r.Add(k)
		}
	}
	return r
}

func Difference(s1, s2 Set) Set {
	r := New()
	for k := range s1.data {
		if _, ok := s2.data[k]; !ok {
			r.Add(k)
		}
	}
	return r
}

func Union(s1, s2 Set) Set {
	r := New()
	for k := range s1.data {
		r.Add(k)
	}
	for k := range s2.data {
		r.Add(k)
	}
	return r
}
