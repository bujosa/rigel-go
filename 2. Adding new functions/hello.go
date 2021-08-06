package hello

import (
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

func Glass() string {
    return quote.Glass()
}

func Proverb() string {
    return quoteV3.Concurrency()
}
