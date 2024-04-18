package exception

type CustomerCreating struct {
}

func (e *CustomerCreating) Error() string {
	return "creating customer failed"
}
