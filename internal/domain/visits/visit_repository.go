package visits

import (
	"context"

	"github.com/vhenzl/url-shortener/internal/domain/links"
)

type VisitRepository interface {
	GetByID(ctx context.Context, id VisitID) (*Visit, error)
	GetAllByLinkID(ctx context.Context, linkID links.LinkID) ([]*Visit, error)
	Add(ctx context.Context, visit *Visit) error
}
