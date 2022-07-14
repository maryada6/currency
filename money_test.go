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

	t.Run("rupee and paise should have same sign ", func(t *testing.T) {
		assert.Panics(t, func() {
			NewMoney(-2, 10)
			NewMoney(2, -10)
		})
	})
}

func TestAdd(t *testing.T) {
	t.Run("add should return money", func(t *testing.T) {
		assert.IsType(t, Money{}, NewMoney(0, 0).Add(NewMoney(0, 0)))
	})

	t.Run("should return money with 5 rupee and 10 paise on addition of 2 rupee and 3 paise with 3 rupee and 7 paise", func(t *testing.T) {
		assert.Equal(t, Money{5, 10}, NewMoney(2, 3).Add(NewMoney(3, 7)))
	})

	t.Run("should return money with -5 rupee and -10 paise on addition of -2 rupee and -3 paise with -3 rupee and -7 paise", func(t *testing.T) {
		assert.Equal(t, Money{-5, -10}, NewMoney(-2, -3).Add(NewMoney(-3, -7)))
	})

	t.Run("should return money with 2 rupee and 0 paise on addition of 3 rupee and 10 paise with -1 rupee and -10 paise", func(t *testing.T) {
		assert.Equal(t, Money{2, 0}, NewMoney(3, 10).Add(NewMoney(-1, -10)))
	})

	t.Run("should return money with 2 rupee and 0 paise on addition of -1 rupee and -10 paise with 3 rupee and 10 paise", func(t *testing.T) {
		assert.Equal(t, Money{2, 0}, NewMoney(-1, -10).Add(NewMoney(3, 10)))
	})

	t.Run("should return money with 151 rupee and 20 paise on addition of 100 rupee and 40 paise with 50 rupee and 80 paise", func(t *testing.T) {
		assert.Equal(t, Money{151, 20}, NewMoney(100, 40).Add(NewMoney(50, 80)))
	})
}
