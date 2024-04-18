package exception

type CustomerCoinShowing struct{}

func (e *CustomerCoinShowing) Error() string {
	return "player coin showing failed"
}
