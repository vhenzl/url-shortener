package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/vhenzl/url-shortener/internal/domain/links"
	"github.com/vhenzl/url-shortener/internal/domain/visits"
)

type VisitRecord struct {
	ID        string    `db:"id"`
	LinkID    string    `db:"link_id"`
	VisitedAt time.Time `db:"visited_at"`
}

type VisitRepository struct {
	db *sqlx.DB
}

// Ensure VisitRepository implements visits.VisitRepository interface
var _ visits.VisitRepository = (*VisitRepository)(nil)

func NewVisitRepository(db *sqlx.DB) *VisitRepository {
	return &VisitRepository{db: db}
}

func (r *VisitRepository) recordToDomain(rec VisitRecord) (*visits.Visit, error) {
	id, err := visits.VisitIDFromString(rec.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid visit ID: %w", err)
	}
	linkID, err := links.LinkIDFromString(rec.LinkID)
	if err != nil {
		return nil, fmt.Errorf("invalid link ID: %w", err)
	}
	return visits.NewVisit(id, linkID, rec.VisitedAt), nil
}

func (r *VisitRepository) domainToRecord(visit *visits.Visit) VisitRecord {
	return VisitRecord{
		ID:        visit.ID().String(),
		LinkID:    visit.LinkID().String(),
		VisitedAt: visit.VisitedAt(),
	}
}

func (r *VisitRepository) GetByID(ctx context.Context, id visits.VisitID) (*visits.Visit, error) {
	var rec VisitRecord
	err := r.db.GetContext(ctx, &rec, "SELECT * FROM visits WHERE id = ?", id.String())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, visits.ErrVisitNotFound
		}
		return nil, fmt.Errorf("get visit by id: %w", err)
	}
	return r.recordToDomain(rec)
}

func (r *VisitRepository) GetAllByLinkID(ctx context.Context, linkID links.LinkID) ([]*visits.Visit, error) {
	var recs []VisitRecord
	err := r.db.SelectContext(ctx, &recs, "SELECT * FROM visits WHERE link_id = ? ORDER BY visited_at DESC", linkID.String())
	if err != nil {
		return nil, fmt.Errorf("get all visits by link id: %w", err)
	}
	visitsList := make([]*visits.Visit, 0, len(recs))
	for _, rec := range recs {
		visit, err := r.recordToDomain(rec)
		if err != nil {
			return nil, err
		}
		visitsList = append(visitsList, visit)
	}
	return visitsList, nil
}

func (r *VisitRepository) Add(ctx context.Context, visit *visits.Visit) error {
	rec := r.domainToRecord(visit)
	_, err := r.db.NamedExecContext(ctx, `
        INSERT INTO visits (id, link_id, visited_at)
        VALUES (:id, :link_id, :visited_at)
    `, &rec)
	if err != nil {
		return fmt.Errorf("add visit: %w", err)
	}
	return nil
}
