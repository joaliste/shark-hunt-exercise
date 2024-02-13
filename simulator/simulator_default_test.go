package simulator_test

import (
	"github.com/stretchr/testify/require"
	"testdoubles/positioner"
	"testdoubles/simulator"
	"testing"
)

func TestCatchSimulatorDefault_CanCatch(t *testing.T) {
	t.Run("prey is hunted", func(t *testing.T) {
		// arrange
		ps := positioner.NewPositionerStub()
		ps.GetLinearDistanceFunc = func(from, to *positioner.Position) (distance float64) {
			distance = 100
			return
		}

		s := simulator.NewCatchSimulatorDefault(100, ps)
		inputHunter := &simulator.Subject{Speed: 10, Position: &positioner.Position{X: 0, Y: 0, Z: 0}}
		inputPrey := &simulator.Subject{Speed: 5, Position: &positioner.Position{X: 100, Y: 0, Z: 0}}

		// act
		value := s.CanCatch(inputHunter, inputPrey)
		// assert
		require.True(t, value)
	})

	t.Run("hunter is slower than prey", func(t *testing.T) {
		// arrange
		ps := positioner.NewPositionerStub()
		ps.GetLinearDistanceFunc = func(from, to *positioner.Position) (distance float64) {
			distance = 100
			return
		}

		s := simulator.NewCatchSimulatorDefault(100, ps)
		inputHunter := &simulator.Subject{Speed: 10, Position: &positioner.Position{X: 0, Y: 0, Z: 0}}
		inputPrey := &simulator.Subject{Speed: 20, Position: &positioner.Position{X: 100, Y: 0, Z: 0}}

		// act
		value := s.CanCatch(inputHunter, inputPrey)
		// assert
		require.False(t, value)
	})

	t.Run("hunter is too far from prey", func(t *testing.T) {
		// arrange
		ps := positioner.NewPositionerStub()
		ps.GetLinearDistanceFunc = func(from, to *positioner.Position) (distance float64) {
			distance = 1000
			return
		}

		s := simulator.NewCatchSimulatorDefault(20, ps)
		inputHunter := &simulator.Subject{Speed: 10, Position: &positioner.Position{X: 0, Y: 0, Z: 0}}
		inputPrey := &simulator.Subject{Speed: 5, Position: &positioner.Position{X: 1000, Y: 0, Z: 0}}

		// act
		value := s.CanCatch(inputHunter, inputPrey)
		// assert
		require.False(t, value)
	})

}
