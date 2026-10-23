package downloaders

import (
	"fmt"
	"os"
	"regexp"
)

var urlSecretPlaceholder = regexp.MustCompile(`\$\{([A-Z][A-Z0-9_]*)\}`)

func expandURLSecrets(rawURL string) (string, error) {
	var missing string
	expanded := urlSecretPlaceholder.ReplaceAllStringFunc(rawURL, func(placeholder string) string {
		name := urlSecretPlaceholder.FindStringSubmatch(placeholder)[1]
		value := os.Getenv(name)
		if value == "" {
			missing = name
		}
		return value
	})
	if missing != "" {
		return "", fmt.Errorf("environment variable %s is not set", missing)
	}
	return expanded, nil
}
