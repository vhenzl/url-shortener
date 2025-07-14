package visits

import "github.com/vhenzl/url-shortener/internal/domain/links"

type VisitRepository interface {
	GetByID(id VisitID) (*Visit, error)
	GetAllByLinkID(linkID links.LinkID) ([]*Visit, error)
	Add(visit *Visit) error
}
