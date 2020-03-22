package singleton

import (
	"github.com/mmcdole/gofeed"
	"sync"
)

// Wrap the gofeed.Parser into singleton
type Parser struct {
	instance *gofeed.Parser
	once     sync.Once
}

func (p *Parser) GetInstance() *gofeed.Parser {
	p.once.Do(func() {
		p.instance = gofeed.NewParser()
	})
	return p.instance
}
