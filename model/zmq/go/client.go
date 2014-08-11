package main

import (
	zmq "github.com/pebbe/zmq3"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting up")
	requester, _ := zmq.NewSocket(zmq.REQ)
	defer requester.Close()
	requester.Connect("tcp://localhost:5555")
	msg := string(`{
  "metrics": [
    { "k": "age",                    "v":  56 },
    { "k": "score",                  "v": 546 },
    { "k": "timeOnHome",             "v":  60 },
    { "k": "timeOnBankSearch",       "v": 400 },
    { "k": "timeOnCCSearch",         "v": 150 },
    { "k": "viewedCCOfferAmex",      "v":   5 },
    { "k": "viewedCCOfferCapOne432", "v":   1 }
  ]
}
`)
	request_nbr := 0
    start := time.Now()
	for  {
		request_nbr++
		requester.Send(msg, 0)
		_, _ = requester.Recv(0)
		if request_nbr % 1000 == 0 {
			elapsed := time.Since(start)
			rate := float64(request_nbr) / elapsed.Seconds();
			fmt.Println("Received: ", request_nbr, elapsed, rate)
		}
	}
}
