package api

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestHandleHTTPError(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusForbidden,
		Body:       io.NopCloser(bytes.NewBufferString(`{"message": "Resource not accessible by integration"}`)),
	}

	err := HandleHTTPError(resp)
	if err == nil || strings.Contains(err.Error(), "gh auth login") {
		t.Errorf("Expected error without auth prompt, got: %v", err)
	}
}