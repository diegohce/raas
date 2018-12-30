package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
)

type fileStorage struct {
	mem   *memStorage
	mutex *sync.RWMutex
}

func (s *fileStorage) init() error {

	s.mem = &memStorage{}
	s.mem.init()

	s.mutex = &sync.RWMutex{}

	content, err := ioutil.ReadFile("./filestorage.json")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(content, s.mem.data); err != nil {
		return err
	}
	return nil
}

func (s *fileStorage) save(sub *subscribeRequest) error {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if err := s.mem.save(sub); err != nil {
		return err
	}

	content, err := json.Marshal(s.mem.data)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile("./filestorage.json", content, os.FileMode(0644)); err != nil {
		return err
	}

	return nil
}

func (s *fileStorage) remove(sub *subscribeRequest) {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.mem.remove(sub)

	content, err := json.Marshal(s.mem.data)
	if err != nil {
		return
	}

	if err := ioutil.WriteFile("./filestorage.json", content, os.FileMode(0644)); err != nil {
		return
	}
}

func (s *fileStorage) subscriptions() []*subscribeRequest {

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.mem.subscriptions()
}
