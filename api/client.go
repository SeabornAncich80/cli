package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type APIError struct {
	Message string `json:"message"`
}

func HandleHTTPError(resp *http.Response) error {
	if resp.StatusCode == http.StatusForbidden {
		body, _ := io.ReadAll(resp.Body)
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err == nil {
			if apiErr.Message == "Resource not accessible by integration" {
				return fmt.Errorf("HTTP 403: %s", apiErr.Message)
			}
		}
		return fmt.Errorf("HTTP 403: Forbidden. Try running 'gh auth login'")
	}
	return nil
}