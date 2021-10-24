package gofish

import "time"

const (
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36"
	Qps = 50
)

var rateLimiter = time.Tick(time.Second / Qps)

type GoFish struct {
	Request *Request
}

func NewGoFish() *GoFish {
	return &GoFish{}
}

func (g *GoFish) Visit() error {
	return g.Request.Do()
}