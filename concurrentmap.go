package main

import (
	"sort"
	"sync"
)

type Reason struct {
	Reason         string `csv:"reason"`
	SpecVoteAction string `csv:"vote_spec_action"`
	KickvoteAction string `csv:"vote_kick_action"`
}
type CSet struct {
	m  map[string]bool
	mu sync.Mutex
}

func NewCSet() CSet {
	m := CSet{
		m: make(map[string]bool),
	}
	return m
}

func (cs *CSet) Add(value string) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.m[value] = true
}

func (cs *CSet) SortedList() []Reason {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	list := make([]string, 0, len(cs.m))
	for value := range cs.m {
		list = append(list, value)
	}

	sort.Sort(sort.StringSlice(list))
	result := make([]Reason, 0, len(list))
	for _, value := range list {
		result = append(result, Reason{
			Reason:         value,
			SpecVoteAction: "unknown",
			KickvoteAction: "unknown",
		})
	}
	return result
}
