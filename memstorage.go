package main

import (
	"fmt"
	"sync"
)

type memStorage struct {
	data  map[string]*subscribeRequest
	mutex *sync.Mutex
}

func (s *memStorage) init() error {
	s.data = make(map[string]*subscribeRequest)
	s.mutex = &sync.Mutex{}
	return nil
}

func (s *memStorage) save(sub *subscribeRequest) error {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	key := fmt.Sprintf("%+v", sub)

	if _, ok := s.data[key]; ok {
		return fmt.Errorf("Subscription %s already exists", key)
	}
	s.data[key] = sub

	return nil
}

func (s *memStorage) remove(sub *subscribeRequest) {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	key := fmt.Sprintf("%+v", sub)

	delete(s.data, key)
}

func (s *memStorage) subscriptions() []*subscribeRequest {

	var rv []*subscribeRequest

	for _, sub := range s.data {
		rv = append(rv, sub)
	}
	return rv
}
