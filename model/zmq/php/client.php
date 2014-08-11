<?php

echo "Making Requests\n";

$context = new ZMQContext();
$requester = new ZMQSocket($context, ZMQ::SOCKET_REQ);
$requester->connect("tcp://localhost:5555");

$json = '{
  "metrics": [
    { "k": "age",                    "v":  56 },
    { "k": "score",                  "v": 546 },
    { "k": "timeOnHome",             "v":  60 },
    { "k": "timeOnBankSearch",       "v": 400 },
    { "k": "timeOnCCSearch",         "v": 150 },
    { "k": "viewedCCOfferAmex",      "v":   5 },
    { "k": "viewedCCOfferCapOne432", "v":   1 }
  ]
}';

//$reply = $requester->send($json)->recv();
//printf ("Received reply: %s\n", $reply);

$ct = 0;
$start = microtime(true);
while (true) {
    $reply = $requester->send($json)->recv();
    if (++$ct % 1000 == 0) {
        echo "$ct - ". number_format($ct / (microtime(true)-$start), 2) ."/s\n";
    }
}
