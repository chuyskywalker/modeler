<?php
/*
* Hello World server
* Binds REP socket to tcp://*:5555
* Expects "Hello" from client, replies with "World"
* @author Ian Barber <ian(dot)barber(at)gmail(dot)com>
*/
echo "Starting\n";
$context = new ZMQContext(1);

// Socket to talk to clients
$responder = new ZMQSocket($context, ZMQ::SOCKET_REP);
$responder->bind("tcp://*:5555");

echo "Bound and waiting\n";

while (true) {
    process($responder);
}

function process(ZMQSocket &$responder) {

    // Wait for next request from client
    $request = $responder->recv();
//    printf ("Received request: [%s]\n", $request);

    // Do some 'work'
    $data = json_decode($request);

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

    // Send reply back to client
    $responder->send($prob);

}