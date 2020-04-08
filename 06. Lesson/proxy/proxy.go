// Imagine we have a server which give us an amount of infected with the coronavirus.
// This server updates an information every 5 seconds. (In fact information updates every 24h,
// but to simulate this process we will take 5 seconds).
// Proxy doest't allow to make many requests to the server and saves up-to-date information.
// Suggest that every 5 seconds number of infected grows.

package proxy

import (
	"math/rand"
	"time"
)

type AmountOfPatients interface {
	Get() int
}

type RealAmountOfPatients struct {
	numberNow int
	lastRequestTime time.Time
}

func (r *RealAmountOfPatients) Get() int {
	r.numberNow += rand.Int() % 300
	r.lastRequestTime = time.Now()
	return r.numberNow
}

type ProxyAmountOfPatients struct {
	realData RealAmountOfPatients
}

func (p *ProxyAmountOfPatients) Get() int {
	if p.realData.numberNow == 0 {
		p.realData.Get()
		return p.realData.numberNow
	} else {
		if time.Now().Sub(p.realData.lastRequestTime).Seconds() < 5 {
			return p.realData.numberNow
		} else {
			p.realData.Get()
			return p.realData.numberNow
		}
	}
}

