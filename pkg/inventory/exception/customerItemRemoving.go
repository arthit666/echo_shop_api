package exception

import "fmt"

type CustomerItemRemoving struct {
	ProductID uint64
}

func (e *CustomerItemRemoving) Error() string {
	return fmt.Sprintf("removing ProductID: %d failed", e.ProductID)
}
