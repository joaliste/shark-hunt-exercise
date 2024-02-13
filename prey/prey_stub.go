package prey

import (
	"testdoubles/positioner"
)

// NewPreyStub creates a new prey stub
func NewPreyStub() *PreyStub {
	return &PreyStub{}
}

// PreyStub is an implementation of the Prey interface
type PreyStub struct {
	// GetSpeedFunc is a replacement of the GetSpeed function
	GetSpeedFunc func() (speed float64)
	// GetPositionFunc is a replacement of the GetPosition function
	GetPositionFunc func() (position *positioner.Position)
}

// GetSpeed returns the speed of the prey
func (ps *PreyStub) GetSpeed() (speed float64) {
	speed = ps.GetSpeedFunc()
	return
}

func (ps *PreyStub) GetPosition() (position *positioner.Position) {
	position = ps.GetPositionFunc()
	return
}
