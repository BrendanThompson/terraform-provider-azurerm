package authorizationruleseventhubs

import "fmt"

const defaultApiVersion = "2021-01-01-preview"

func userAgent() string {
	return fmt.Sprintf("pandora/authorizationruleseventhubs/%s", defaultApiVersion)
}
