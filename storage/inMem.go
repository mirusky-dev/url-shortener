package storage

import (
	"errors"

	"github.com/mirusky-dev/url-shortener/util"
)

// InMem is a simple in memory Storage
type InMem struct {
	db map[string]string
}

// Save create a new data
func (s *InMem) Save(url string) (string, error) {
	id := util.Random(6)
	s.db[id] = url
	return id, nil
}

// Get return informed data from id
func (s *InMem) Get(id string) (string, error) {
	v, ok := s.db[id]
	if !ok {
		return "", errors.New("Not found")
	}
	return v, nil
}

// NewInMem return a new in-men storage
func NewInMem() Storage {
	return &InMem{
		make(map[string]string),
	}
}
