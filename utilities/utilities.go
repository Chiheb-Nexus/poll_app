package utilities

import "fmt"

// FormatAPIVersion is a helper function to format API endpoints with APP version
func FormatAPIVersion(endpoint string, APIVersion string) string {
	return fmt.Sprintf("/api/%s/%s", APIVersion, endpoint)
}
