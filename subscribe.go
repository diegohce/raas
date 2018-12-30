package main

import (
	"github.com/diegohce/easyrabbit"
	"log"
)

type subscribeRequest struct {
	RmqURI       string `json:"rmq_uri"`
	QueueName    string `json:"queue_name"`
	CallbackURL  string `json:"callback_url"`
	stopConsumer chan bool
}

func (sr *subscribeRequest) subscribe() error {

	erc, err := easyrabbit.New(sr.RmqURI)
	if err != nil {
		return err
	}

	sr.stopConsumer, err = erc.ConsumeWithCallback(sr.QueueName, "raas", sr.consume)
	if err != nil {
		return err
	}

	if err := storage.save(sr); err != nil {
		log.Println("ERROR subscribe::", err, "saving subscription")
	}

	return nil
}
