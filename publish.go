package main

import (
	"github.com/diegohce/easyrabbit"
)

type publishRequest struct {
	RmqURI       string `json:"rmq_uri"`
	ExchangeName string `json:"exchange_name"`
	RoutingKey   string `json:"routing_key"`
	Body         string `json:"body"`
}

func (pr *publishRequest) publish() error {

	erc, err := easyrabbit.New(pr.RmqURI)
	if err != nil {
		return err
	}
	defer erc.Close()

	if err := erc.Publish(pr.ExchangeName,
		pr.RoutingKey, []byte(pr.Body)); err != nil {
			return err
	}
	return nil
}
