package simulator

// NewCatchSimulatorMock creates a new CatchSimulatorDefault
func NewCatchSimulatorMock() *CatchSimulatorMock {
	return &CatchSimulatorMock{}
}

// CatchSimulatorMock is a default implementation of CatchSimulator
type CatchSimulatorMock struct {
	// CanCatchFunc externalize the CanCatch method
	CanCatchFunc func(hunter, prey *Subject) (ok bool)

	// Observer
	Calls struct {
		// CanCatch is the number of times the CanCatch method has been called
		CanCatch int
	}
}

// CanCatch returns true if the hunter can catch the prey
func (m *CatchSimulatorMock) CanCatch(hunter, prey *Subject) (ok bool) {
	ok = m.CanCatchFunc(hunter, prey)
	m.Calls.CanCatch++
	return ok
}
