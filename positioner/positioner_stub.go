package positioner

// NewPositionerStub returns a new NewPositionerStub instance
func NewPositionerStub() (positioner *PositionerStub) {
	positioner = &PositionerStub{}
	return
}

// PositionerStub is a struct that represents a default positioner stub
type PositionerStub struct {
	// GetLinearDistance function to use a GetLinearDistance real function
	GetLinearDistanceFunc func(from, to *Position) (linearDistance float64)
}

// GetLinearDistance returns the linear distance between 2 positions (in meters)
func (ps *PositionerStub) GetLinearDistance(from, to *Position) (linearDistance float64) {
	linearDistance = ps.GetLinearDistanceFunc(from, to)
	return
}
