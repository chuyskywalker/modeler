<?php

echo "Starting\n";

$address = '127.0.0.1';
$port    = 4000;

if (($sock = socket_create(AF_INET, SOCK_STREAM, SOL_TCP)) === false) {
    echo "socket_create() failed: reason: " . socket_strerror(socket_last_error()) . "\n";
}

if (socket_bind($sock, $address, $port) === false) {
    echo "socket_bind() failed: reason: " . socket_strerror(socket_last_error($sock)) . "\n";
}

if (socket_listen($sock, 100) === false) {
    echo "socket_listen() failed: reason: " . socket_strerror(socket_last_error($sock)) . "\n";
}

while(1) {

    if (($msgsock = socket_accept($sock)) === false) {
        echo "socket_accept() failed: reason: " . socket_strerror(socket_last_error($sock)) . "\n";
        break;
    }

    if (false === ($buf = socket_read($msgsock, 2048))) {
        echo "socket_read() failed: reason: " . socket_strerror(socket_last_error($msgsock)) . "\n";
        break;
    }

    if (!$buf = trim($buf)) {
        continue;
    }

    $data = json_decode($buf);

    $metricsSimple = [];
    foreach ($data->metrics as $m) {
        $metricsSimple[$m->k] = $m->v;
    }

    $prob = json_encode(['probability' => (
              ($metricsSimple['age'] * .4)
            + ($metricsSimple['score'] * .8)
            + ($metricsSimple['timeOnHome'] * .04)
            + ($metricsSimple['timeOnBankSearch'] * .1)
            + ($metricsSimple['timeOnCCSearch'] * .2)
            + ($metricsSimple['viewedCCOfferAmex'] * .4)
            + ($metricsSimple['viewedCCOfferCapOne432'] * .4)
        )]);

    socket_write($msgsock, $prob, strlen($prob));
    socket_close($msgsock);

};

socket_close($sock);
