package exception

import "fmt"

type ProductNotFound struct {
	ProductID uint64
}

func (e *ProductNotFound) Error() string {
	return fmt.Sprintf("itemID: %d was not found", e.ProductID)
}
