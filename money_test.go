package currency

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMoney(t *testing.T) {
	t.Run("should be able to initialize money with rupee and paise", func(t *testing.T) {
		assert.NotPanics(t, func() {
			NewMoney(1, 0)
		})
	})

	t.Run("new money should return money", func(t *testing.T) {
		assert.IsType(t, Money{}, NewMoney(0, 0))
	})

	t.Run("should convert excess paise to rupee", func(t *testing.T) {
		assert.Equal(t, NewMoney(10, 10), NewMoney(0, 1010))
	})
}
