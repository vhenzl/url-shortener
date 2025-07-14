package links

import (
	"time"

	"github.com/google/uuid"
)

type URL string
// Slug is a user-facing link identifier.
type Slug string
// LinkID is a type alias for link UUIDs.
type LinkID uuid.UUID

// Link represents a named URL mapping.
type Link struct {
	ID        LinkID
	Slug      Slug
	TargetURL URL
	CreatedAt time.Time
	UpdatedAt time.Time
}
