package util

import (
	"sync"
)

type Set struct {
	m map[string] bool
	sync.RWMutex
}

func New() *Set {
	return &Set{
		m: map[string]bool{},
	}
}

func (s *Set) Add(item string) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

func (s *Set) remove(item string) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}

func (s *Set) Has(item string) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

func (s *Set) Len() int {
	return len(s.List())
}

func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[string]bool{}
}

func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

func (s *Set) List() []string {
	s.RLock()
	defer s.RUnlock()
	var list []string
	for item := range s.m {
		list = append(list, item)
	}
	return list
}


