package exception

import "fmt"

type ProductArchiving struct {
	ProductID uint64
}

func (e *ProductArchiving) Error() string {
	return fmt.Sprintf("archiving product id: %d failed", e.ProductID)
}
