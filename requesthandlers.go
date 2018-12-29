package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
)

func publishHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")

	preq := &publishRequest{}

	body, err := unmarshalBody(r, preq)
	if err != nil {
		log.Println("ERROR on publishHandler::", err, ":: Request body:", string(body))
		w.WriteHeader(400)
		fmt.Fprintf(w, "{\"result\":\"%s\"}", err)
		return
	}

	if err := preq.publish(); err != nil {
		log.Println("ERROR on publishHandler::", err, ":: Request body:", string(body))
		w.WriteHeader(400)
		fmt.Fprintf(w, "{\"result\":\"%s\"}", err)
		return
	}

	fmt.Fprintf(w, "{\"result\":\"ok\"}")
}

func subscribeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")

	sreq := &subscribeRequest{}

	body, err := unmarshalBody(r, sreq)
	if err != nil {
		log.Println("ERROR on subscribeHandler::", err, ":: Request body:", string(body))
		w.WriteHeader(400)
		fmt.Fprintf(w, "{\"result\":\"%s\"}", err)
		return
	}

	if err := sreq.subscribe(); err != nil {
		log.Println("ERROR on subscribeHandler::", err, ":: Request body:", string(body))
		w.WriteHeader(400)
		fmt.Fprintf(w, "{\"result\":\"%s\"}", err)
		return
	}

	fmt.Fprintf(w, "{\"result\":\"ok\"}")
}

func unmarshalBody(r *http.Request, target interface{}) ([]byte, error) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, target); err != nil {
		return body, err
	}

	return body, nil
}
