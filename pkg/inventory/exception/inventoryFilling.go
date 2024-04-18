package exception

import "fmt"

type InventoryFilling struct {
	CustomerID string
	ProductID  uint64
}

func (e *InventoryFilling) Error() string {
	return fmt.Sprintf("inventory filling for CustomerID: %s and ProductID: %d failed", e.CustomerID, e.ProductID)
}
