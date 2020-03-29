package observer

import "testing"

func TestSomeObserved_Notify(t *testing.T) {
	var obs1 Obs1
	var obs2 Obs2
	var observed *SomeObserved
	observed = new(SomeObserved)

	observed.AddObserver(obs1)
	observed.AddObserver(obs2)

	out := observed.Notify()

	if out != "12" {
		t.Errorf("Error testing notify chain")
	}
}

func TestSomeObserved_Remove(t *testing.T) {
	var obs1 Obs1
	var obs2 Obs2
	var observed *SomeObserved
	observed = new(SomeObserved)

	observed.AddObserver(obs1)
	observed.AddObserver(obs2)

	observed.RemoveObserver(obs1)
	out := observed.Notify()

	if out != "2" {
		t.Errorf("Error testing remove")
	}
}
