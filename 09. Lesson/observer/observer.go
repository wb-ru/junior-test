package observer

type Publisher interface {
	AddMarket(observer Observer)
	ChangePrice(price int)
	Notify()
}

type Observer interface {
	Update(price int)
	getMarketName() string
}

type item struct {
	observers []Observer
	name      string
	price     int
}

func newItem(name string) *item {
	return &item{
		name: name,
	}
}

func (s *item) AddMarket(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *item) ChangePrice(price int) {
	s.price = price
	s.Notify()
}

func (s *item) Notify() {
	for _, observer := range s.observers {
		observer.Update(s.price)
	}
}

type Market struct {
	name  string
	price int
}

func (s *Market) Update(price int) {
	s.price = price
}

func (s *Market) getMarketName() string {
	return s.name
}
