package domain

import (
	"time"

	"github.com/google/uuid"
)

type URL string
type Slug string
type LinkID uuid.UUID

// Link represents a named URL mapping.
type Link struct {
	ID        LinkID // Internal unique ID
	Slug      string // User-facing link identifier
	TargetURL URL    // The destination URL
	CreatedAt time.Time
	UpdatedAt time.Time
}
