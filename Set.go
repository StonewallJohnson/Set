package set

import "fmt"

type Set interface {
	union(*Set) Set
	intersection(*Set) Set
	difference(*Set) Set
	subset(*Set) bool
}

type StaticIntSet struct {
	//This will implement a Set
	size_     int
	container map[int]bool
}

//Psuedo default constructor
func CreateEmptySet() StaticIntSet {
	return StaticIntSet{size_: 0, container: make(map[int]bool)}
}

/**
Adds the given int into the set
returns true if successful, false if not
**/
func (s *StaticIntSet) add(x int) bool {
	if s.contains(x) {
		return false
	}
	s.container[x] = true
	s.size_++
	return true
}

/**
Removes the given int from the set
returns true if the element was removed, false if it was not present
**/
func (s *StaticIntSet) remove(x int) bool {
	if s.contains(x) {
		delete(s.container, x)
		s.size_--
		return true
	}
	return false
}

//Returns whether or not the given int is in the set
func (s *StaticIntSet) contains(x int) bool {
	ok := s.container[x]
	if ok {
		return true
	} else {
		return false
	}
}

func (s *StaticIntSet) size() int {
	return s.size_
}

func testMembership(o *Set) StaticIntSet {
	otherSet := *o
	other, ok := otherSet.(StaticIntSet)
	if !ok {
		//inappropriate type in other
		//don't know what to do here yet
		//end program or something?
		fmt.Printf("Cannot do set operations on StaticIntSet and %T", *o)
		return other
	}
	return other
}

func (s StaticIntSet) union(other *Set) Set {
	otherStaticIntSet := testMembership(other)

	toReturn := CreateEmptySet()

	for key := range otherStaticIntSet.container {
		//add all elements of other
		toReturn.add(key)
	}

	for key := range s.container {
		//add all elements of s
		toReturn.add(key)
	}

	var newSet Set = toReturn
	return newSet

}

func (s StaticIntSet) difference(other *Set) Set {
	otherSet := testMembership(other)

	toReturn := CreateEmptySet()
	for key := range s.container {
		//for all keys in s
		if !otherSet.contains(key) {
			//if the key is not in otherset, put into toReturn
			toReturn.add(key)
		}
	}

	var set Set = toReturn
	return set
}

func (s StaticIntSet) intersection(other *Set) Set {
	otherSet := testMembership(other)

	toReturn := CreateEmptySet()
	for key := range s.container {
		//for every key in s
		if otherSet.contains(key) {
			//if in other set, add to toReturn
			toReturn.add(key)
		}
	}

	for key := range otherSet.container {
		//for every key in other
		if s.contains(key) {
			//if in s, add to toReturn
			//duplicate adding is not allowed
			toReturn.add(key)
		}
	}

	var set Set = toReturn
	return set
}

func (s StaticIntSet) subset(other *Set) bool {
	otherSet := testMembership(other)

	for key := range otherSet.container {
		//every key in otherSet
		if !s.contains(key) {
			//the key is not in s
			return false
		}
	}
	return true
}
