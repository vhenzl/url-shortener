package links

import "context"

type LinkRepository interface {
	GetByID(ctx context.Context, id LinkID) (*Link, error)
	GetBySlug(ctx context.Context, slug Slug) (*Link, error)
	Add(ctx context.Context, link *Link) error
	Update(ctx context.Context, link *Link) error
	Remove(ctx context.Context, link *Link) error
}
