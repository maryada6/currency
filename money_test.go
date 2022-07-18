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
		assert.IsType(t, float64(0), NewMoney(0, 0).Add(NewMoney(0, 0)))
	})

	t.Run("should return money with 5 rupee and 10 paise on addition of 2 rupee and 3 paise with 3 rupee and 7 paise", func(t *testing.T) {
		assert.Equal(t, 5.10, NewMoney(2, 3).Add(NewMoney(3, 7)))
	})

	t.Run("should return money with -5 rupee and -10 paise on addition of -2 rupee and -3 paise with -3 rupee and -7 paise", func(t *testing.T) {
		assert.Equal(t, -5.10, NewMoney(-2, -3).Add(NewMoney(-3, -7)))
	})

	t.Run("should return money with 2 rupee and 0 paise on addition of 3 rupee and 10 paise with -1 rupee and -10 paise", func(t *testing.T) {
		assert.Equal(t, 2.00, NewMoney(3, 10).Add(NewMoney(-1, -10)))
	})

	t.Run("should return money with 2 rupee and 0 paise on addition of -1 rupee and -10 paise with 3 rupee and 10 paise", func(t *testing.T) {
		assert.Equal(t, 2.00, NewMoney(-1, -10).Add(NewMoney(3, 10)))
	})

	t.Run("should return money with 151 rupee and 20 paise on addition of 100 rupee and 40 paise with 50 rupee and 80 paise", func(t *testing.T) {
		assert.Equal(t, 151.20, NewMoney(100, 40).Add(NewMoney(50, 80)))
	})
}

func TestAmount(t *testing.T) {
	t.Run("should return float value", func(t *testing.T) {
		assert.IsType(t, float64(0), NewMoney(100, 100).Amount())
	})

	t.Run("should return 0.00 as amount for 0 rupee and 0 paise", func(t *testing.T) {
		assert.Equal(t, 0.00, NewMoney(0, 0).Amount())
	})

	t.Run("should return 1.00 as amount for 1 rupee and 0 paise", func(t *testing.T) {
		assert.Equal(t, 1.00, NewMoney(1, 0).Amount())
	})

	t.Run("should return 100.00 as amount for 0 rupee and 10000 paise", func(t *testing.T) {
		assert.Equal(t, 100.00, NewMoney(0, 10000).Amount())
	})

	t.Run("should return -100.00 as amount for 0 rupee and -10000 paise", func(t *testing.T) {
		assert.Equal(t, -100.00, NewMoney(0, -10000).Amount())
	})

	t.Run("should return -1.00 as amount for -1 rupee and 0 paise", func(t *testing.T) {
		assert.Equal(t, -1.00, NewMoney(-1, 0).Amount())
	})

	t.Run("should return 100.10 as amount for 100 rupee and 10 paise", func(t *testing.T) {
		assert.Equal(t, 100.10, NewMoney(100, 10).Amount())
	})

	t.Run("should return -200.00 as amount for -100 rupee and -10000 paise", func(t *testing.T) {
		assert.Equal(t, -200.00, NewMoney(-100, -10000).Amount())
	})
}

func TestEqual(t *testing.T) {
	t.Run("should return boolean value", func(t *testing.T) {
		moneyOne := NewMoney(1, 10)
		moneyTwo := NewMoney(0, 110)
		assert.IsType(t, true, moneyOne.Equal(moneyTwo))
		assert.IsType(t, true, moneyTwo.Equal(moneyOne))
	})

	t.Run("should return true value for equal money", func(t *testing.T) {
		moneyOne := NewMoney(1, 10)
		moneyTwo := NewMoney(1, 10)
		assert.IsType(t, true, moneyOne.Equal(moneyTwo))
		assert.IsType(t, true, moneyTwo.Equal(moneyOne))
	})

	t.Run("should return false value for different money", func(t *testing.T) {
		moneyOne := NewMoney(1, 10)
		moneyTwo := NewMoney(9, 110)
		assert.IsType(t, true, moneyOne.Equal(moneyTwo))
		assert.IsType(t, true, moneyTwo.Equal(moneyOne))
	})

	t.Run("should return false value for different negative money", func(t *testing.T) {
		assert.Equal(t, false, NewMoney(-110, -10).Equal(NewMoney(-99, -110)))
		assert.Equal(t, false, NewMoney(-99, -110).Equal(NewMoney(-110, -10)))
	})

	t.Run("should return true value for equal negative money", func(t *testing.T) {
		moneyOne := NewMoney(-99, -10)
		moneyTwo := NewMoney(-199, -110)
		assert.IsType(t, true, moneyOne.Equal(moneyTwo))
		assert.IsType(t, true, moneyTwo.Equal(moneyOne))
	})

	t.Run("should return false value while comparing a negative money and a positive money", func(t *testing.T) {
		moneyOne := NewMoney(1, 10)
		moneyTwo := NewMoney(-1, -10)
		assert.IsType(t, true, moneyOne.Equal(moneyTwo))
		assert.IsType(t, true, moneyTwo.Equal(moneyOne))
	})

	t.Run("should follow transitive property", func(t *testing.T) {
		moneyOne := NewMoney(2, 10)
		moneyTwo := NewMoney(0, 210)
		moneyThree := NewMoney(1, 110)
		assert.IsType(t, true, moneyOne.Equal(moneyTwo))
		assert.IsType(t, true, moneyTwo.Equal(moneyThree))
		assert.IsType(t, true, moneyOne.Equal(moneyThree))
	})
}

func TestSubtract(t *testing.T) {
	t.Run("add should return money", func(t *testing.T) {
		assert.IsType(t, float64(0), NewMoney(0, 0).Subtract(NewMoney(0, 0)))
	})

	t.Run("should return money with -1 rupee and 0 paise on subtraction of 2 rupee and 7 paise with 3 rupee and 7 paise", func(t *testing.T) {
		assert.Equal(t, -1.00, NewMoney(2, 7).Subtract(NewMoney(3, 7)))
	})

	t.Run("should return money with 1 rupee and 0 paise on subtraction of -2 rupee and -7 paise with -3 rupee and -7 paise", func(t *testing.T) {
		assert.Equal(t, 1.00, NewMoney(-2, -7).Subtract(NewMoney(-3, -7)))
	})

	t.Run("should return money with 5.14 rupee and 0 paise on subtraction of 2 rupee and 7 paise with -3 rupee and -7 paise", func(t *testing.T) {
		assert.Equal(t, 5.14, NewMoney(2, 7).Subtract(NewMoney(-3, -7)))
	})

	t.Run("should return money with -5.14 rupee and 0 paise on subtraction of -2 rupee and -7 paise with 3 rupee and 7 paise", func(t *testing.T) {
		assert.Equal(t, -5.14, NewMoney(-2, -7).Subtract(NewMoney(3, 7)))
	})

	t.Run("should return money with 5.14 rupee and 0 paise on subtraction of 5 rupee and 14 paise with 0 rupee and 0 paise", func(t *testing.T) {
		assert.Equal(t, 5.14, NewMoney(5, 14).Subtract(NewMoney(0, 0)))
	})

	t.Run("should return money with -5.14 rupee and 0 paise on subtraction of -5 rupee and -14 paise with 0 rupee and 0 paise", func(t *testing.T) {
		assert.Equal(t, -5.14, NewMoney(-5, -14).Subtract(NewMoney(0, 0)))
	})
}
