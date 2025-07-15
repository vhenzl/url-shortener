package infrastructure

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/vhenzl/url-shortener/internal/domain/links"
	"github.com/vhenzl/url-shortener/internal/domain/visits"
)

type VisitRecord struct {
	ID        visits.VisitID `db:"id"`
	LinkID    links.LinkID   `db:"link_id"`
	VisitedAt time.Time      `db:"visited_at"`
}

type VisitRepository struct {
	db *sqlx.DB
}

func NewVisitRepository(db *sqlx.DB) *VisitRepository {
	return &VisitRepository{db: db}
}

func (r *VisitRepository) recordToDomain(rec VisitRecord) *visits.Visit {
	return visits.NewVisit(rec.ID, rec.LinkID, rec.VisitedAt)
}

func (r *VisitRepository) domainToRecord(visit *visits.Visit) VisitRecord {
	return VisitRecord{
		ID:        visit.ID(),
		LinkID:    visit.LinkID(),
		VisitedAt: visit.VisitedAt(),
	}
}

func (r *VisitRepository) GetByID(id visits.VisitID) (*visits.Visit, error) {
	var rec VisitRecord
	err := r.db.Get(&rec, "SELECT * FROM visits WHERE id = ?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, visits.ErrVisitNotFound
		}
		return nil, fmt.Errorf("get visit by id: %w", err)
	}
	return r.recordToDomain(rec), nil
}

func (r *VisitRepository) GetAllByLinkID(linkID links.LinkID) ([]*visits.Visit, error) {
	var recs []VisitRecord
	err := r.db.Select(&recs, "SELECT * FROM visits WHERE link_id = ? ORDER BY visited_at DESC", linkID)
	if err != nil {
		return nil, fmt.Errorf("get all visits by link id: %w", err)
	}
	visitsList := make([]*visits.Visit, 0, len(recs))
	for _, rec := range recs {
		visitsList = append(visitsList, r.recordToDomain(rec))
	}
	return visitsList, nil
}

func (r *VisitRepository) Add(visit *visits.Visit) error {
	rec := r.domainToRecord(visit)
	_, err := r.db.NamedExec(`
        INSERT INTO visits (id, link_id, visited_at)
        VALUES (:id, :link_id, :visited_at)
    `, &rec)
	if err != nil {
		return fmt.Errorf("add visit: %w", err)
	}
	return nil
}
