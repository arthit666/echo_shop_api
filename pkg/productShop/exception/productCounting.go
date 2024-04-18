package exception

type ProductCounting struct{}

func (e *ProductCounting) Error() string {
	return "item counting failed"
}
