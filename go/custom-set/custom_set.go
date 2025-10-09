package stringset

import "strings"

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set map[string]struct{}

func New() Set {
	return make(Set)
}

func NewFromSlice(l []string) Set {
	set := New()
	for _, s := range l {
		set[s] = struct{}{}
	}
	return set
}

func (s Set) String() string {
	var sb strings.Builder
	sb.WriteRune('{')
	i := 0
	for k := range s {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString("\"" + k + "\"")
		i++
	}
	sb.WriteRune('}')
	return sb.String()
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s[elem]
	return ok
}

func (s Set) Add(elem string) {
	s[elem] = struct{}{}
}

func Subset(s1, s2 Set) bool {
	if len(s1) == 0 {
		return true
	}
	for k := range s1 {
		if !s2.Has(k) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for k := range s1 {
		if s2.Has(k) {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}
	for k := range s1 {
		if !s2.Has(k) {
			return false
		}
	}
	return true
}

func Intersection(s1, s2 Set) Set {
	inter := New()
	for k := range s1 {
		if s2.Has(k) {
			inter.Add(k)
		}
	}
	return inter
}

func Difference(s1, s2 Set) Set {
	inter := New()
	for k := range s1 {
		if !s2.Has(k) {
			inter.Add(k)
		}
	}
	return inter
}

func Union(s1, s2 Set) Set {
	inter := New()
	for k := range s1 {
		inter.Add(k)
	}
	for k := range s2 {
		inter.Add(k)
	}
	return inter
}
