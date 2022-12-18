package validation

import (
	"strings"
)

const specialChars = `\!@#$%&*()_+\-=\[\]{};:"\\|,.<>\/?`

func IsSanitized(value string) bool {
	return !strings.ContainsAny(value, specialChars)
}
