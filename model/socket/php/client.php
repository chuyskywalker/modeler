<?php

echo "Making Requests\n";

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

function doCon() {
    global $json;

    $socket = socket_create(AF_INET, SOCK_STREAM, SOL_TCP);
    if ($socket === false) {
        echo "socket_create() failed: reason: " . socket_strerror(socket_last_error()) . "\n";
    }

    $address = '127.0.0.1';
    $port    = 4000;

    $result = socket_connect($socket, $address, $port);
    if ($result === false) {
        echo "socket_connect() failed.\nReason: ($result) " . socket_strerror(socket_last_error($socket)) . "\n";
    }

    socket_write($socket, $json, strlen($json));

    $out = socket_read($socket, 2048);

    socket_close($socket);

    return $out;
}

//echo doCon() . "\n";

$ct = 0;
$start = microtime(true);
while (true) {
    doCon();
    if (++$ct % 1000 == 0) {
        echo "$ct - ". number_format($ct / (microtime(true)-$start), 2) ."/s\n";
    }
}
