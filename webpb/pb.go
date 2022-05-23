package webpb

import (
	"sync"
)

func NewProgressBar(val int, max int) (pb *ProgressBar) {
	pb = &ProgressBar{
		val: val,
		max: max,
	}

	return
}

type ProgressBar struct {
	mu      sync.Mutex
	updated bool

	val int
	max int
}

func (p *ProgressBar) Inc() {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.val < p.max {
		p.val++
		p.updated = true
	}
}

func (p *ProgressBar) Add(delta int) {
	if delta == 0 {
		return
	}

	p.mu.Lock()
	defer p.mu.Unlock()
	if p.val+delta <= p.max {
		p.val += delta
		p.updated = true
	} else {
		if p.val != p.max {
			p.val = p.max
			p.updated = true
		}
	}
}

// Updated - Если счетчик (p.val) был увеличен, то при первом вызове вернет - true, а при всех последующих - false
func (p *ProgressBar) Updated() bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.updated {
		p.updated = false
		return true
	}
	return false
}
