package main

import (
	"github.com/diegohce/easyrabbit"
)

type subscribeRequest struct {
	RmqURI      string `json:"rmq_uri"`
	QueueName   string `json:"queue_name"`
	CallbackURL string `json:"callback_url"`
}

func (sr *subscribeRequest) subscribe() error {

	erc, err := easyrabbit.New(sr.RmqURI)
	if err != nil {
		return err
	}

	if err := erc.ConsumeWithCallback(sr.QueueName, "raas", sr.consume); err != nil {
		return err
	}

	return nil
}
