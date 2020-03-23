package chainOfResponsibility

import "testing"

func TestChainOne(t *testing.T) {
	// setting up the chain
	h1 := new(FirstHandler)
	h2 := new(SecondHandler)
	h3 := new(ThirdHandler)

	h1.next = h2
	h2.next = h3
	handled := new(Handled)
	h1.Exec(handled)
	want := "123"
	if handled.info != want {
		t.Errorf("Got %s, want %s", handled.info, want)
	}
}

func TestChainTwo(t *testing.T) {
	// setting up the chain
	h1 := new(FirstHandler)
	h2 := new(SecondHandler)
	h3 := new(ThirdHandler)

	h1.next = h3
	h2.next = h2
	handled := new(Handled)
	h1.Exec(handled)
	want := "132"
	if handled.info != want {
		t.Errorf("Got %s, want %s", handled.info, want)
	}
}
