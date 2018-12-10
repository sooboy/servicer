package servicer

import (
	"sync"
)

// Paralleler 并发执行Servicer
type Paralleler interface {
	Push(...Servicer)
	Run()
}

// Parallel implements Paralleler
type Parallel struct {
	items []Servicer
	wait  sync.WaitGroup
}

// Push  implements Paralleler
func (p *Parallel) Push(items ...Servicer) {
	if p.items == nil {
		p.items = make([]Servicer, 0)
	}
	p.items = append(p.items, items...)
}

// Run implements Paralleler
func (p *Parallel) Run() {
	for i := 0; i < len(p.items); i++ {
		p.wait.Add(1)
		go func(num int) {
			p.items[num].Run()
			p.wait.Done()
		}(i)
	}
	p.wait.Wait()
}
