package visits

import (
	"time"

	"github.com/google/uuid"
	"github.com/vhenzl/url-shortener/internal/domain/links"
)

// VisitID is a type alias for visit UUIDs.
type VisitID uuid.UUID

// Visit represents a single access to a link.
type Visit struct {
	ID        VisitID
	LinkID    links.LinkID
	VisitedAt time.Time
}
