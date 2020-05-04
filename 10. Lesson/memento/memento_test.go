package memento

import (
	"testing"
)

func TestMomento(t *testing.T) {

	originator := new(Actual)
	caretaker := new(Caretaker)
	originator.version = "1.0"
	caretaker.t = originator.Update()
	originator.version = "1.1"
	originator.Rollback(caretaker.t)

	if originator.version != "1.0" {
		t.Errorf("Expect State to %s, but %s", originator.version, "1.0")
	}
}
