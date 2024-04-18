package exception

import "fmt"

type ProductEditing struct {
	ItemID uint64
}

func (e *ProductEditing) Error() string {
	return fmt.Sprintf("editing product id: %d failed", e.ItemID)
}
