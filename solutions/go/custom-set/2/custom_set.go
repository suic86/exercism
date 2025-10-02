package stringset

import (
	"fmt"
	"strings"
)

type Set map[string]bool

func New() Set {
	return make(Set)
}

func NewFromSlice(l []string) Set {
	s := New()
	for _, k := range l {
		s[k] = true
	}
	return s
}

func (s Set) String() string {
	l := []string{}
	for k := range s {
		l = append(l, `"`+k+`"`)
	}
	return fmt.Sprintf("{%v}", strings.Join(l, ", "))
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s[elem]
	return ok
}

func (s Set) Add(elem string) {
	s[elem] = true
}

func Subset(s1, s2 Set) bool {
	for k := range s1 {
		if _, ok := s2[k]; !ok {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for k := range s1 {
		if _, ok := s2[k]; ok {
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
	for k := range s1 {
		if _, ok := s2[k]; ok {
			r.Add(k)
		}
	}
	return r
}

func Difference(s1, s2 Set) Set {
	r := New()
	for k := range s1 {
		if _, ok := s2[k]; !ok {
			r.Add(k)
		}
	}
	return r
}

func Union(s1, s2 Set) Set {
	r := New()
	for k := range s1 {
		r.Add(k)
	}
	for k := range s2 {
		r.Add(k)
	}
	return r
}
