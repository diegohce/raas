package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type fileStorage struct {
	data map[string]*subscribeRequest
}

func (s *fileStorage) init() error {
	s.data = make(map[string]*subscribeRequest)

	content, err := ioutil.ReadFile("./filestorage.json")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(content, s.data); err != nil {
		return err
	}
	return nil
}

func (s *fileStorage) save(sub *subscribeRequest) error {

	key := fmt.Sprintf("%+v", sub)

	if _, ok := s.data[key]; ok {
		return fmt.Errorf("Subscription %s already exists", key)
	}
	s.data[key] = sub

	content, err := json.Marshal(s.data)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile("./filestorage.json", content, os.FileMode(0644)); err != nil {
		return err
	}

	return nil
}

func (s *fileStorage) remove(sub *subscribeRequest) {

	key := fmt.Sprintf("%+v", sub)

	delete(s.data, key)

	content, err := json.Marshal(s.data)
	if err != nil {
		return
	}

	if err := ioutil.WriteFile("./filestorage.json", content, os.FileMode(0644)); err != nil {
		return
	}
}

func (s *fileStorage) subscriptions() []*subscribeRequest {

	var rv []*subscribeRequest

	for _, sub := range s.data {
		rv = append(rv, sub)
	}
	return rv
}
