package main

import (
	"fmt"
	"sort"
	"strings"
)

// New types can be defined based on existing types.
// In this case we create a new type for storing strings
// as we want to implement a String() method for it that
// does not exist for the builtin type. Note that both types
// are distinct even though their values may look similar.
type MyString string

func (s MyString) String() string {
	return string(s)
}

type Address struct {
	name   string
	street string
	zip    int
}

func (a Address) String() string {
	return fmt.Sprintf("%v | %v | %v", a.name, a.street, a.zip)
}

// Type constraints allow defining expectations on generic types
// in a central, declared way. We can define that the specific
// type for the generic type Element needs to be comparable
// as well as provide a method for returning a string representation.
type Element interface {
	comparable
	String() string
}

type Set[E Element] struct {
	internalMap map[E]bool
}

func NewSet[E Element](elements ...E) Set[E] {
	internalMap := map[E]bool{}

	for _, element := range elements {
		internalMap[element] = true
	}

	return Set[E]{internalMap}
}

func (s Set[E]) Size() int {
	return len(s.internalMap)
}

func (s *Set[E]) Add(element E) {
	s.internalMap[element] = true
}

func (s *Set[E]) Remove(element E) {
	delete(s.internalMap, element)
}

func (s Set[E]) Contains(element E) bool {
	// Since we only store true values for keys that are existing,
	// we can just return the value. Missing key will result in default
	// value which is false for bool types.
	return s.internalMap[element]
}

func (s Set[E]) Slice() []E {
	slice := make([]E, 0, len(s.internalMap))

	for element := range s.internalMap {
		slice = append(slice, element)
	}

	return slice
}

// Due to our type constraint we can rely on the fact that every element
// type has a String() method that we can use for creating a complete
// string representation of this set. Elements are ordered ascending.
func (s Set[E]) String() string {
	slice := []string{}
	for element := range s.internalMap {
		slice = append(slice, element.String())
	}

	sort.Strings(slice)

	return strings.Join(slice, ", ")
}
