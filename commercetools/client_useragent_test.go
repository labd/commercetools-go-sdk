package commercetools_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/stretchr/testify/assert"
)

func TestUserAgents(t *testing.T) {
	testCases := []struct {
		cfg               *commercetools.ClientConfig
		expectedUserAgent string
	}{
		{
			cfg: &commercetools.ClientConfig{
				LibraryName:    "terraform-provider-commercetools",
				LibraryVersion: "0.1",
				ContactURL:     "https://example.com",
				ContactEmail:   "test@example.org",
			},
			expectedUserAgent: fmt.Sprintf(
				"commercetools-go-sdk/%s Go/%s (%s; %s) terraform-provider-commercetools/0.1 (+https://example.com; +test@example.org)",
				commercetools.Version, runtime.Version(), runtime.GOOS, runtime.GOARCH),
		},
		{
			cfg: &commercetools.ClientConfig{
				ContactURL:   "https://example.com",
				ContactEmail: "test@example.org",
			},
			expectedUserAgent: fmt.Sprintf(
				"commercetools-go-sdk/%s Go/%s (%s; %s) (+https://example.com; +test@example.org)",
				commercetools.Version, runtime.Version(), runtime.GOOS, runtime.GOARCH),
		},
		{
			cfg: &commercetools.ClientConfig{
				LibraryName:    "terraform-provider-commercetools",
				LibraryVersion: "0.1",
				ContactEmail:   "test@example.org",
			},
			expectedUserAgent: fmt.Sprintf(
				"commercetools-go-sdk/%s Go/%s (%s; %s) terraform-provider-commercetools/0.1 (+test@example.org)",
				commercetools.Version, runtime.Version(), runtime.GOOS, runtime.GOARCH),
		},
		{
			cfg: &commercetools.ClientConfig{
				LibraryName:  "terraform-provider-commercetools",
				ContactURL:   "https://example.com",
				ContactEmail: "test@example.org",
			},
			expectedUserAgent: fmt.Sprintf(
				"commercetools-go-sdk/%s Go/%s (%s; %s) terraform-provider-commercetools (+https://example.com; +test@example.org)",
				commercetools.Version, runtime.Version(), runtime.GOOS, runtime.GOARCH),
		},
	}

	for _, tC := range testCases {
		t.Run("Test user agent", func(t *testing.T) {
			userAgent := commercetools.GetUserAgent(tC.cfg)
			assert.Equal(t, tC.expectedUserAgent, userAgent)
		})
	}
}
