package observer

type Observer interface {
	Handle() string
}

type Observed interface {
	AddObserver(obs Observer)
	RemoveObserver(obs Observer)
	Notify() string
}

var Payload = []string{"1", "2"}

type Obs1 struct {
}

func (p Obs1) Handle() string {
	return Payload[0]
}

type Obs2 struct{}

func (p Obs2) Handle() string {
	return Payload[1]
}

type SomeObserved struct {
	storage map[Observer]bool
}

func (p *SomeObserved) AddObserver(obs Observer) {
	if p.storage == nil {
		p.storage = make(map[Observer]bool)
	}
	p.storage[obs] = true
}

func (p *SomeObserved) RemoveObserver(obs Observer) {
	delete(p.storage, obs)
}

func (p *SomeObserved) Notify() string {
	res := ""
	for key, _ := range p.storage {
		res += key.Handle()
	}
	return res
}
