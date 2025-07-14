package links

import (
	"time"

	"github.com/google/uuid"
)

type URL string
type Slug string
type LinkID uuid.UUID

// Link represents a named URL mapping.
type Link struct {
	id        LinkID
	slug      Slug
	targetURL URL
	createdAt time.Time
	updatedAt time.Time
}

func NewLink(id LinkID, slug Slug, target URL, createdAt, updatedAt time.Time) *Link {
	return &Link{
		id:        id,
		slug:      slug,
		targetURL: target,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (l *Link) ID() LinkID           { return l.id }
func (l *Link) Slug() Slug           { return l.slug }
func (l *Link) TargetURL() URL       { return l.targetURL }
func (l *Link) CreatedAt() time.Time { return l.createdAt }
func (l *Link) UpdatedAt() time.Time { return l.updatedAt }
