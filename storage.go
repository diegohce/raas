package main

import (
	"log"
)

type storer interface {
	init() error
	save(*subscribeRequest) error
	remove(*subscribeRequest)
	subscriptions() []*subscribeRequest
}

var (
	storers map[string]storer
	storage storer
)

const (
	storeName = "file" //"mem"
)

func init() {

	storers = map[string]storer{
		"mem":  &memStorage{},
		"file": &fileStorage{},
	}

	storage = storers[storeName]
	storage.init()

	for _, s := range storage.subscriptions() {
		if err := s.subscribe(); err != nil {
			log.Println("subscribe::init", err)
		}
	}
}
