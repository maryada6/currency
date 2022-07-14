package currency

type Money struct {
	rupee int64
	paise int64
}

func NewMoney(rupee, paise int64) Money {
	if (rupee < 0 && paise > 0) || (rupee > 0 && paise < 0) {
		panic("rupee and paise should have same sign")
	}
	if paise >= 100 || paise < 0 {
		rupee += paise / 100
		paise = paise % 100
	}

	return Money{rupee, paise}
}

func (moneyOne Money) Add(moneyTwo Money) Money {
	paise := moneyOne.rupee*100 + moneyTwo.rupee*100 + moneyOne.paise + moneyTwo.paise

	return NewMoney(0, paise)
}
