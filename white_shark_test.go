package hunt_test

import (
	"github.com/stretchr/testify/require"
	hunt "testdoubles"
	"testing"
)

// Tests for the WhiteShark implementation - Hunt method
func TestWhiteSharkHunt(t *testing.T) {
	t.Run("case 1: white shark hunts successfully", func(t *testing.T) {
		// Arrange
		s := hunt.NewWhiteShark(
			true,
			false,
			200)
		tuna := hunt.NewTuna("Tuna 1", 100)
		// Act
		err := s.Hunt(tuna)
		// Assert
		require.Nil(t, err)
		require.True(t, s.Tired)
		require.False(t, s.Hungry)
	})

	t.Run("case 2: white shark is not hungry", func(t *testing.T) {
		// Arrange
		s := hunt.NewWhiteShark(
			false,
			false,
			200)
		tuna := hunt.NewTuna("Tuna 1", 200)
		// Act
		err := s.Hunt(tuna)
		// Assert
		require.ErrorIs(t, err, hunt.ErrSharkIsNotHungry, "shark is not hungry")
	})

	t.Run("case 3: white shark is tired", func(t *testing.T) {
		// Arrange
		s := hunt.NewWhiteShark(
			true,
			true,
			200)
		tuna := hunt.NewTuna("Tuna 1", 200)
		// Act
		err := s.Hunt(tuna)
		// Assert
		require.ErrorIs(t, err, hunt.ErrSharkIsTired, "shark is tired")
	})

	t.Run("case 4: white shark is slower than the tuna", func(t *testing.T) {
		// Arrange
		s := hunt.NewWhiteShark(
			true,
			false,
			200)
		tuna := hunt.NewTuna("Tuna 1", 300)
		// Act
		err := s.Hunt(tuna)
		// Assert
		require.ErrorIs(t, err, hunt.ErrSharkIsSlower, "shark is too slow")
	})

	t.Run("case 5: tuna is nil", func(t *testing.T) {
		// Arrange
		s := hunt.NewWhiteShark(
			true,
			false,
			200)
		// Act
		err := s.Hunt(nil)
		// Assert
		require.ErrorIs(t, err, hunt.ErrTunaIsNil, "tuna is nil")
	})
}
