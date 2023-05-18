package api

import "fmt"

type ErrorResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf(`[%s]:%s`, e.Name, e.Message)
}
