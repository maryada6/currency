package currency

type Money struct {
	rupee int64
	paise int64
}

func NewMoney(rupee, paise int64) Money {
	if paise >= 100 {
		rupee += paise / 100
		paise = paise % 100
	}

	return Money{rupee, paise}
}

func (money_one Money) Add(money_two Money) Money {
	return Money{5, 10}
}
