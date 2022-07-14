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

func (moneyOne Money) Add(moneyTwo Money) float64 {
	paise := moneyOne.paise + moneyTwo.paise

	return NewMoney(0, paise).Amount()
}

func (moneyOne Money) Amount() float64 {
	return float64(moneyOne.paise) / 100
}
