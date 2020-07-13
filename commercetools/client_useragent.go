package commercetools

import (
	"fmt"
	"runtime"
	"strings"
)

// GetUserAgent determines the user agent for all HTTP requests.
func GetUserAgent(cfg *ClientConfig) string {
	baseInfo := fmt.Sprintf("commercetools-go-sdk/%s", Version)
	systemInfo := fmt.Sprintf("Go/%s (%s; %s)", runtime.Version(), runtime.GOOS, runtime.GOARCH)

	libraryInfo := ""
	if cfg.LibraryName != "" && cfg.LibraryVersion == "" {
		libraryInfo = cfg.LibraryName
	} else if cfg.LibraryName != "" && cfg.LibraryVersion != "" {
		libraryInfo = fmt.Sprintf("%s/%s", cfg.LibraryName, cfg.LibraryVersion)
	}
	contactInfo := ""
	if cfg.ContactURL != "" && cfg.ContactEmail == "" {
		contactInfo = fmt.Sprintf("(+%s)", cfg.ContactURL)
	} else if cfg.ContactURL == "" && cfg.ContactEmail != "" {
		contactInfo = fmt.Sprintf("(+%s)", cfg.ContactEmail)
	} else if cfg.ContactURL != "" && cfg.ContactEmail != "" {
		contactInfo = fmt.Sprintf("(+%s; +%s)", cfg.ContactURL, cfg.ContactEmail)
	}

	s := []string{
		baseInfo,
		systemInfo,
	}
	if libraryInfo != "" {
		s = append(s, libraryInfo)
	}
	if contactInfo != "" {
		s = append(s, contactInfo)
	}

	return strings.Join(s, " ")
}
