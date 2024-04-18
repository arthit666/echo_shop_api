package exception

type ProductCreating struct{}

func (e *ProductCreating) Error() string {
	return "creating product failed"
}
