package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func (sr *subscribeRequest) consume(body []byte) error {

	cbMessage := struct {
		QueueName string `json:"queue_name"`
		Body      string `json:"body"`
	}{
		QueueName: sr.QueueName,
		Body:      string(body),
	}

	jMessage, err := json.Marshal(cbMessage)
	if err != nil {
		log.Println("ERROR on consumer for:", sr.CallbackUrl, sr.QueueName, "::", err)
		return err
	}

	res, err := http.Post(sr.CallbackUrl, "application/json", bytes.NewReader(jMessage))
	if err != nil {
		log.Println("ERROR on consumer for:", sr.CallbackUrl, sr.QueueName, "::", err, "body:", string(jMessage))
		return err
	}
	defer res.Body.Close()

	return nil
}

