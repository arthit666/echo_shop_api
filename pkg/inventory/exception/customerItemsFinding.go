package exception

import "fmt"

type CustomerItemsFinding struct {
	CustomerID string
}

func (e *CustomerItemsFinding) Error() string {
	return fmt.Sprintf("finding customer product for customerID: %s failed", e.CustomerID)
}
