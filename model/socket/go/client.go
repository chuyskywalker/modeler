package main

import (
	"fmt"
	"time"
	"net"
	"bufio"
)

func main() {
	fmt.Println("Starting up")

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
	tlen := 0
    start := time.Now()
	for  {
		request_nbr++

		c, err := net.Dial("tcp", "127.0.0.1:4001")
		if err != nil {
			fmt.Println(err)
		}

//		fmt.Println("Sending: ", msg)
		c.Write([]byte(msg))

		data, err := bufio.NewReader(c).ReadString('\n')
		tlen += len(data)
//		fmt.Println("Received: ", status)

		c.Close()

		if request_nbr % 1000 == 0 {
			elapsed := time.Since(start)
			rate := float64(request_nbr) / elapsed.Seconds();
			fmt.Println("Received: ", request_nbr, elapsed, rate)
		}
	}
}
