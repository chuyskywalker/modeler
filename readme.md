
Send over:

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


Server Computes:
```
  (age * .4)
+ (score * .8)
+ (timeOnHome * .04)
+ (timeOnBankSearch * .1)
+ (timeOnCCSearch * .2)
+ (viewedCCOfferAmex * .4)
+ (viewedCCOfferCapOne432 * .4)
```

And returns:
```
{ "probability": 534 }
```