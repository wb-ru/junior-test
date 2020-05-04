package memento

type Actual struct {
	version string
}

func (a *Actual) Update() *Testing {
	return &Testing{version: a.version}
}

func (a *Actual) Rollback(t *Testing) {
	a.version = t.ReturnState()
}

type Testing struct {
	version string
}

func (t *Testing) ReturnState() string {
	return t.version
}

type Caretaker struct {
	t *Testing
}
