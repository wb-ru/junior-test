package memento

type Memento struct {
	state string
}

func (p *Memento) getState() string {
	return p.state
}

type Originator struct {
	state string
}

func (p *Originator) AcceptState(m *Memento) {
	p.state = m.state
}

func (p *Originator) GetState() *Memento {
	m := Memento{state: p.state}
	return &m
}

type CareTaker struct {
	holder []*Memento
}

func (p *CareTaker) Push(m *Memento) {
	p.holder = append(p.holder, m)
}

func (p *CareTaker) Pop() *Memento {
	var ret *Memento
	if len(p.holder) == 0 {
		return nil
	}
	p.holder, ret = p.holder[:len(p.holder)-1], p.holder[len(p.holder)-1]
	return ret
}
