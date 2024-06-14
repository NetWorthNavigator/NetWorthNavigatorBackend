package utils

import "strings"

func NormalizeEmail(email string) string {
	// Split the email into local part and domain part
	parts := strings.Split(email, "@")
	if len(parts) == 2 {
		localPart := parts[0]
		domainPart := strings.ToLower(parts[1]) // Convert domain part to lowercase

		// Normalize only for Gmail addresses
		if domainPart == "gmail.com" {
			localPart = strings.ReplaceAll(localPart, ".", "") // Remove dots from local part
		}
		return localPart + "@" + domainPart
	}
	return strings.ToLower(email) // Default case: just convert to lowercase
}
