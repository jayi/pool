package pool

import "sync"

type Pool struct {
	connPool chan bool
	wg *sync.WaitGroup
}

func NewPool(maxConnCount int) (*Pool) {
	pool := new(Pool)
	pool.wg = new(sync.WaitGroup)
	if maxConnCount < 1 {
		maxConnCount = 1
	}
	pool.connPool = make(chan bool, maxConnCount)
	for i := 0; i < maxConnCount; i++ {
		pool.connPool <- true
	}
	return pool
}

func (pool *Pool)Get() {
	pool.wg.Add(1)
	<-pool.connPool
}

func (pool *Pool)Put() {
	pool.wg.Done()
	pool.connPool <- true
}

func (pool *Pool)Wait() {
	pool.wg.Wait()
}
