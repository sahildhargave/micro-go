package util

const (
	USD = "USD"
	EUR = "EUR"
	RUP = "RUP"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, RUP:
		return true
	}
	return false
}
