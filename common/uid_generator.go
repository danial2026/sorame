package common

import (
	"strings"

	"github.com/google/uuid"
)

// generateLinkUid generates a unique UID to use as LinkUid
func GenerateLinkUid() string {
	// Generate a new UUID
	u := uuid.New().String()

	// Extract only alphanumeric characters
	alphanumeric := strings.ReplaceAll(u, "-", "") // Remove hyphens
	alphanumeric = alphanumeric[:8]                // Take the first 8 characters

	return alphanumeric
}