package storage

import (
	"github.com/MirableOne/word-of-wisdom/pkg/hashcash"
	"sync"
)

type InMemoryMapStorage struct {
	container sync.Map
}

func NewInMemoryMapStorage() *InMemoryMapStorage {
	return &InMemoryMapStorage{}
}

func (s *InMemoryMapStorage) Exists(header string) bool {
	_, ok := s.container.Load(header)
	return ok
}

func (s *InMemoryMapStorage) Push(header string) {
	s.container.Store(header, 1)
}

var _ hashcash.Storage = (*InMemoryMapStorage)(nil)
