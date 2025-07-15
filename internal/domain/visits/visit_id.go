package visits

import (
	"github.com/google/uuid"
)

type VisitID uuid.UUID

func NewVisitID() VisitID {
	u, err := uuid.NewV7()
	if err != nil {
		panic("failed to generate VisitID: " + err.Error())
	}
	return VisitID(u)
}

func VisitIDFromString(s string) (VisitID, error) {
	u, err := uuid.Parse(s)
	return VisitID(u), err
}

func (id VisitID) UUID() uuid.UUID {
	return uuid.UUID(id)
}

func (id VisitID) String() string {
	return id.UUID().String()
}
