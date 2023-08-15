package auth_test

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {

	got, err := auth.GetAPIKey(http.Header{})
	if (got) != "" {
		t.Error("There was a error")
	}
	if (err) == nil {
		t.Error("There was a error")
	}
}
