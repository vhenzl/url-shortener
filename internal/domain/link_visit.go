package domain

import (
	"time"

	"github.com/google/uuid"
)

type LinkVisitID uuid.UUID

// LinkVisit represents a single access to a link.
type LinkVisit struct {
	ID        LinkVisitID
	LinkID    LinkID
	VisitedAt time.Time
}
