package exception

type ProductListing struct{}

func (e *ProductListing) Error() string {
	return "product listing failed"
}
