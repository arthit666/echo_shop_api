package exception

import "fmt"

type CustomerNotFound struct {
	CustomID string
}

func (e *CustomerNotFound) Error() string {
	return fmt.Sprintf("CutomerID: %s not found", e.CustomID)
}
