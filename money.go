package currency

const conversion = 100

type Money struct {
	paise int64
}

func NewMoney(rupee, paise int64) Money {
	if (rupee < 0 && paise > 0) || (rupee > 0 && paise < 0) {
		panic("rupee and paise should have same sign")
	}
	paise = rupee*conversion + paise

	return Money{paise}
}

func (money Money) Add(moneyTwo Money) float64 {
	paise := money.paise + moneyTwo.paise

	return NewMoney(0, paise).Amount()
}

func (money Money) Amount() float64 {
	return float64(money.paise) / conversion
}

func (money Money) Equal(moneyTwo Money) bool {
	return money.paise == moneyTwo.paise
}

func (money Money) Subtract(moneyTwo Money) float64 {
	moneyTwoFlipSign := NewMoney(0, -moneyTwo.paise)

	return money.Add(moneyTwoFlipSign)
}
