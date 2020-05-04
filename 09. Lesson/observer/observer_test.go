package observer

func ExampleObserver() {
	publisher := newItem("Bananas")
	observer1 := &Market{name: "Rynok"}
	observer2 := &Market{name: "Ne-rynok"}
	publisher.AddMarket(observer1)
	publisher.AddMarket(observer2)
	publisher.ChangePrice(60)
	publisher.Notify()
}
