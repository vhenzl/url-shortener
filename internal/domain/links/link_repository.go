package links

type LinkRepository interface {
	GetByID(id LinkID) (*Link, error)
	GetBySlug(slug Slug) (*Link, error)
	Add(link *Link) error
	Update(link *Link) error
	Remove(link *Link) error
}
