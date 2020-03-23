package chainOfResponsibility

type Handled struct {
	info string
}
type Handler interface {
	Exec(h *Handled)
}

type FirstHandler struct {
	next Handler
}

func (p *FirstHandler) Exec(h *Handled) {
	h.info += "1"
	if p.next != nil {
		p.next.Exec(h)
	}
}

type SecondHandler struct {
	next Handler
}

func (p *SecondHandler) Exec(h *Handled) {
	h.info += "2"
	if p.next != nil {
		p.next.Exec(h)
	}
}

type ThirdHandler struct {
	next Handler
}

func (p *ThirdHandler) Exec(h *Handled) {
	h.info += "3"
	if p.next != nil {
		p.next.Exec(h)
	}
}
