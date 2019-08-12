package token

import (
	"encoding/hex"

	"github.com/gofrs/uuid"
)

// New creates a new UUID-based token encoded in hex.
func New() string {
	return hex.EncodeToString(uuid.Must(uuid.NewV4()).Bytes())
}

// NewShort creates a new UUID-based short token encoded in hex.
func NewShort() string {
	return New()[:16]
}
