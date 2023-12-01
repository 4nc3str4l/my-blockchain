package utils

import "encoding/json"

func JsonStatus(message string) string {
	m, _ := json.Marshal(struct {
		Message string `json:"message"`
	}{
		Message: message,
	})
	return string(m)
}
