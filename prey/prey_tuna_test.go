package prey_test

import (
	"github.com/stretchr/testify/require"
	"testdoubles/prey"
	"testing"
)

func TestTuna_GetSpeed(t *testing.T) {
	t.Run("speed is 0.0", func(t *testing.T) {
		// arrange
		tuna := prey.NewTuna(0.0, nil)
		expectedResult := 0.0
		// act
		speed := tuna.GetSpeed()
		// assert
		require.Equal(t, expectedResult, speed)
	})

	t.Run("speed greater than 0", func(t *testing.T) {
		// arrange
		expectedSpeed := 200.0
		tuna := prey.NewTuna(expectedSpeed, nil)

		// act
		speed := tuna.GetSpeed()
		// assert
		require.Equal(t, expectedSpeed, speed)
	})
}
