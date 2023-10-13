package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Set struct {
	data [10]string
	size int
}

type interfaceSet interface {
	add(value string) bool
	contains(value string) bool
	delete(value string) bool
	getSize() int
	printSet()
}

func (s *Set) add(value string) bool {
	if s.contains(value) {
		return false
	}
	s.data[s.size] = value
	s.size++
	return true
}

func (s *Set) contains(value string) bool {
	for _, datum := range s.data {
		if datum == value {
			return true
		}
	}
	return false
}

func (s *Set) indexOf(value string) int {
	for i := 0; i < s.size; i++ {
		if value == s.data[i] {
			return i
		}
	}
	return -1
}

func (s *Set) delete(value string) bool {
	if s.contains(value) {
		for i := s.indexOf(value); i <= s.size; i++ {
			s.data[i] = s.data[i+1]
		}
		s.size--
		return true
	}
	return false
}

func (s *Set) getSize() int {
	return s.size
}

func (s *Set) printSet() {
	fmt.Println(s.data)
}

func TestAdd(t *testing.T) {
	var obj interfaceSet
	obj = &Set{}
	assert.True(t, obj.add("nathan"))
	assert.True(t, obj.add("fernando"))
	assert.True(t, obj.add("lumban"))
	assert.True(t, obj.add("tobing"))

	assert.False(t, obj.add("lumban"))
	assert.False(t, obj.add("tobing"))
}

func TestContains(t *testing.T) {
	obj := interfaceSet(&Set{})
	obj.add("nathan")
	obj.add("fernando")

	assert.True(t, obj.contains("nathan"))
	assert.True(t, obj.contains("fernando"))

	assert.False(t, obj.contains("lumban"))
	assert.False(t, obj.contains("tobing"))
}

func TestDelete(t *testing.T) {
	obj := interfaceSet(&Set{})
	obj.add("value1")
	obj.add("value2")
	obj.add("value3")
	obj.add("value4")
	obj.add("value5")
	obj.printSet()
	fmt.Println(obj.getSize())
	assert.True(t, obj.delete("value4"))
	fmt.Println(obj.getSize())
	obj.printSet()

	obj.add("nathan")
}
