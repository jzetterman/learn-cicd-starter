package auth_test

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"github.com/stretchr/testify/assert"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("valid header", func(t *testing.T) {
		// build request with Authorization header
		// call auth.GetAPIKey(r.Header)
		// check key and error
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		req.Header.Set("Authorization", "ApiKey abc123")
		key, err := auth.GetAPIKey(req.Header)
		if err != nil {
			t.Errorf("unexpected error (valid header): %v", err)
		}
		assert.NoError(t, err)
		assert.Equal(t, "abc123", key)
	})
	t.Run("missing key", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		req.Header.Set("Authorization", "ApiKey")
		key, err := auth.GetAPIKey(req.Header)

		assert.Error(t, err)
		assert.Empty(t, key)
	})
}
