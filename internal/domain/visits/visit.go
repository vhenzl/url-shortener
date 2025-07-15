package visits

import (
	"errors"
	"time"

	"github.com/vhenzl/url-shortener/internal/domain/links"
)

var ErrVisitNotFound = errors.New("visit not found")

// Visit represents a single access to a link.
type Visit struct {
	id        VisitID
	linkID    links.LinkID
	visitedAt time.Time
}

func NewVisit(id VisitID, linkID links.LinkID, visitedAt time.Time) *Visit {
	return &Visit{
		id:        id,
		linkID:    linkID,
		visitedAt: visitedAt,
	}
}

func (v *Visit) ID() VisitID          { return v.id }
func (v *Visit) LinkID() links.LinkID { return v.linkID }
func (v *Visit) VisitedAt() time.Time { return v.visitedAt }
