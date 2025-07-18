// SPDX-License-Identifier: GPL-3.0-or-later
package validators

import (
	"strings"
)

// IsValidLanguageCode validates ISO 639-1 language codes and that we support the code.
func IsValidLanguageCode(lang string) bool {
	allowedLanguages := map[string]bool{
		"en": true,
		"fr": true,
	}

	// Basic validation: exactly 2 lowercase letters.
	if len(lang) != 2 {
		return false
	}

	return allowedLanguages[lang]
}

// SanitizeLanguageCode ensures language code is safe for table name construction.
func SanitizeLanguageCode(lang string) string {
	if !IsValidLanguageCode(lang) {
		return ""
	}
	// Convert to uppercase for table name prefix.
	return strings.ToUpper(lang)
}

// IsLanguageSupported checks if a language exists in the provided list.
func IsLanguageSupported(lang string, availableLanguages []string) bool {
	for _, availableLang := range availableLanguages {
		if availableLang == lang {
			return true
		}
	}
	return false
}
