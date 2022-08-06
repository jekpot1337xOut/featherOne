package utils

import (
	"sync"
)

var thread = 15

// TODO
// to test the difference location of wg
// in pool struct or global
var wg sync.WaitGroup

type Pool struct {
	targetIPList []string
	resultIPList []*Response
	reqChan      chan *Request
}

func NewPool(targetIPList []string) *Pool {
	return &Pool{
		targetIPList: targetIPList,
		reqChan:      make(chan *Request, 100),
	}
}

// Start entry of download pool
func (p *Pool) Start() {
	for _, targetIP := range p.targetIPList {
		p.reqChan <- NewRequest("GET", targetIP, nil)
	}
	close(p.reqChan)
	for i := 0; i < thread; i++ {
		wg.Add(1)
		go p.doGoroutine()
	}
	wg.Wait()
}

// goroutine jobs
func (p *Pool) doGoroutine() {
	defer wg.Done()
	for req := range p.reqChan {
		resp := req.Do()
		colorOut(resp)
		p.resultIPList = append(p.resultIPList, resp)
	}
}
