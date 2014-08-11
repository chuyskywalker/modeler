package main

import (
    "net"
    "log"
    "fmt"
//    "io"
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
    l, err := net.Listen("tcp", "localhost:4001")
    if err != nil {
        log.Fatal(err)
    }
    for {
        c, err := l.Accept()
        if err != nil {
            log.Fatal(err)
        }
        go serve(c)
    }
}

func serve(c net.Conn) {

    j := json.NewDecoder(c)

    var request Request
    err := j.Decode(&request)
    if err != nil {
        fmt.Println("error:", err)
    }
//    fmt.Printf("rec: %+v\n", request)

    // turn metrics into lookup map
    var lookup map[string]float32
    lookup = make(map[string]float32)
    for _, value := range request.Metrics {
        lookup[value.K] = float32(value.V);
    }
//    fmt.Printf("%+v\n", lookup)

    p := ((lookup["age"] * 0.4) + (lookup["score"] * 0.8) + (lookup["timeOnHome"] * 0.04) + (lookup["timeOnBankSearch"] * 0.1) + (lookup["timeOnCCSearch"] * 0.2) + (lookup["viewedCCOfferAmex"] * 0.4) + (lookup["viewedCCOfferCapOne432"] * 0.4))
    prob := Response{Probability: p}

//    fmt.Printf("%+v\n", prob)

    b, err := json.Marshal(prob)
    if err != nil {
        fmt.Println("error:", err)
    }

    // send it back!
    c.Write(b)

	c.Close()
}
