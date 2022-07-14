package currency

type Money struct {
	paise int64
}

func NewMoney(rupee, paise int64) Money {
	if (rupee < 0 && paise > 0) || (rupee > 0 && paise < 0) {
		panic("rupee and paise should have same sign")
	}
	paise = rupee*100 + paise

	return Money{paise}
}

func (money Money) Add(moneyTwo Money) float64 {
	paise := money.paise + moneyTwo.paise

	return NewMoney(0, paise).Amount()
}

func (money Money) Amount() float64 {
	return float64(money.paise) / 100
}

func (money Money) Equal(moneyTwo Money) bool {
	if money.paise == moneyTwo.paise {
		return true
	}
	
	return false
}
