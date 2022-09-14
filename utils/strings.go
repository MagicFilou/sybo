package utils

import (
	"github.com/google/uuid"
)

// IsValidUUID: convenience function to check if a string is a valid uuid.
func IsValidUUID(u string) bool {

	_, err := uuid.Parse(u)
	return err == nil
}
