package singleton

import (
	"testing"
)

func TestParser_GetInstance(t *testing.T) {
	Parser := new(Parser)
	other := Parser.GetInstance()
	another := Parser.GetInstance()
	if other != another {
		t.Errorf("Instances do not match!")
	}
}
