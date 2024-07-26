package main

import "sync"

type SafeMap struct {
	sync.RWMutex
	m map[int]int
}

func (m SafeMap) Get(key int) (int, bool) {
	m.Lock()
	defer m.Unlock()
	v, exists := m.m[key]
	return v, exists
}

func (m SafeMap) Set(key, value int) {
	m.Lock()
	defer m.Unlock()
	m.m[key] = value
}

