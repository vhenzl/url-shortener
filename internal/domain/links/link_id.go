package links

import (
	"github.com/google/uuid"
)

type LinkID uuid.UUID

func NewLinkID() LinkID {
	u, err := uuid.NewV7()
	if err != nil {
		panic("failed to generate LinkID: " + err.Error())
	}
	return LinkID(u)
}

func LinkIDFromString(s string) (LinkID, error) {
	u, err := uuid.Parse(s)
	return LinkID(u), err
}

func (id LinkID) UUID() uuid.UUID {
	return uuid.UUID(id)
}

func (id LinkID) String() string {
	return id.UUID().String()
}
