package exception

import "fmt"

type ProductNotEnough struct {
	ProductID uint64
}

func (e *ProductNotEnough) Error() string {
	return fmt.Sprintf("ProductID: %d is not enough", e.ProductID)
}
