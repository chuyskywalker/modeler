# Modeler

Wanted to develop a quick POC to demonstrate "send a set of measurements, 
calculate against a data model, return results" and bench PHP against Go.

Found that Go was unbelieveable slower than PHP. Seeking advice as to what
I surely must be doing wrong with the Go code.

## Concept

### Clients

Clients will send over a data packet like this:

```json
{
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
```

This data format is a bit more complex than a straight map of string->values, ie:

```json
{
  "age": 56,
  "score": 546
  ...etc...
}
```

However, the slightly more complex format allows us to easily add additional
information should it be needed. For example, perhaps each metric needs an
added, supplimental `weight`. Or we need to add some kind of top level
information that's not a metric factor.

### Servers

Upon receipt of the data, the server will decode and then compute:

```
  (age * .4)
+ (score * .8)
+ (timeOnHome * .04)
+ (timeOnBankSearch * .1)
+ (timeOnCCSearch * .2)
+ (viewedCCOfferAmex * .4)
+ (viewedCCOfferCapOne432 * .4)
```

Which, with our static input (as above) would produce 534 which should then
be encoded to a JSON object and returned as such:

```json
{ "probability": 534 }
```

## Running

First, make sure you have Vagrant and VirtualBox installed, then run 
`vagrant up && vagrant ssh` to get into the system. The `up` process
will install golang and php along with Zeromq. The ZMQ library for go
is cached in this repo (probably not the right way, but for this demo,
it'll do.)

Once logged in you can run each of the clients and servers in 
combination (they are cross language compatible) and measure the results.

```bash
cd /vagrant/modeler/go
go build srv.go
go build client.go
cd ..

# Start PHP server:
php php/srv.php &

# Test clients against PHP srv:
php php/client.php # End when satisfied
./go/client        # End when satisfied

# Kill PHP server
killall -9 php

# Start go server:
./go/srv

# Again, test clients
php php/client.php # End when satisfied
./go/client        # End when satisfied
```

## Results

_Add grains of salt, comes from my own machine, but should maintain 
like ratios:_

| client | srv | Per second (@10k) |
| --- | --- | ------- |
| php | php | 1,294/s |
| go  | php |   915/s |
| php | go  |   872/s |
| go  | go  |   672/s |

> wat...

I would expect that list to be exactly reversed -- I've got to believe 
something I am doing in Go is very wrong.