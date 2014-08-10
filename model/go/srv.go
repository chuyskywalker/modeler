package main

import (
	zmq "github.com/pebbe/zmq3"
	"fmt"
	"encoding/json"
)

type Metric struct {
	K string `json:"k"`
	V int `json:"v"`
}

type Request struct {
	Metrics []Metric
}

type Response struct {
	Probability float32 `json:"probability"`
}

func main() {
	responder, _ := zmq.NewSocket(zmq.REP)
	defer responder.Close()
	responder.Bind("tcp://*:5555")

	for {
		msg, _ := responder.Recv(0)
//		fmt.Println("Received ", msg)

		var request Request
		err := json.Unmarshal([]byte(msg), &request)
		if err != nil {
			fmt.Println("error:", err)
		}
//		fmt.Printf("%+v\n", request)

		// turn metrics into lookup map
		var lookup map[string]float32
		lookup = make(map[string]float32)
		for _, value := range request.Metrics {
			lookup[value.K] = float32(value.V);
		}

//		fmt.Printf("%+v\n", lookup)

		p := ((lookup["age"] * 0.4) + (lookup["score"] * 0.8) + (lookup["timeOnHome"] * 0.04) + (lookup["timeOnBankSearch"] * 0.1) + (lookup["timeOnCCSearch"] * 0.2) + (lookup["viewedCCOfferAmex"] * 0.4) + (lookup["viewedCCOfferCapOne432"] * 0.4))
		prob := Response{Probability: p}

//		fmt.Printf("%+v\n", prob)

		b, err := json.Marshal(prob)
		if err != nil {
			fmt.Println("error:", err)
		}

		// send it back!
		responder.Send(string(b), 0)

//		responder.Send("hi", 0)
	}

}
