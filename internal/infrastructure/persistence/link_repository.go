package infrastructure

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/vhenzl/url-shortener/internal/domain/links"
)

type LinkRecord struct {
	ID        string    `db:"id"`
	Slug      string    `db:"slug"`
	TargetURL string    `db:"target_url"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type LinkRepository struct {
	db *sqlx.DB
}

// Ensure LinkRepository implements links.LinkRepository interface
var _ links.LinkRepository = (*LinkRepository)(nil)

func NewLinkRepository(db *sqlx.DB) *LinkRepository {
	return &LinkRepository{db: db}
}

func (r *LinkRepository) recordToDomain(rec LinkRecord) (*links.Link, error) {
	id, err := links.LinkIDFromString(rec.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid link ID: %w", err)
	}
	return links.NewLink(
		id,
		links.Slug(rec.Slug),
		links.URL(rec.TargetURL),
		rec.CreatedAt,
		rec.UpdatedAt,
	), nil
}

func (r *LinkRepository) domainToRecord(link *links.Link) LinkRecord {
	return LinkRecord{
		ID:        link.ID().String(),
		Slug:      string(link.Slug()),
		TargetURL: string(link.TargetURL()),
		CreatedAt: link.CreatedAt(),
		UpdatedAt: link.UpdatedAt(),
	}
}

func (r *LinkRepository) GetByID(id links.LinkID) (*links.Link, error) {
	var rec LinkRecord
	err := r.db.Get(&rec, "SELECT * FROM links WHERE id = ?", id.String())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, links.ErrLinkNotFound
		}
		return nil, fmt.Errorf("get by id: %w", err)
	}
	return r.recordToDomain(rec)
}

func (r *LinkRepository) GetBySlug(slug links.Slug) (*links.Link, error) {
	var rec LinkRecord
	err := r.db.Get(&rec, "SELECT * FROM links WHERE slug = ?", slug)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, links.ErrLinkNotFound
		}
		return nil, fmt.Errorf("get by slug: %w", err)
	}
	return r.recordToDomain(rec)
}

func (r *LinkRepository) Add(link *links.Link) error {
	rec := r.domainToRecord(link)
	_, err := r.db.NamedExec(`
        INSERT INTO links (id, slug, target_url, created_at, updated_at)
        VALUES (:id, :slug, :target_url, :created_at, :updated_at)
    `, &rec)
	if err != nil {
		return fmt.Errorf("add link: %w", err)
	}
	return nil
}

func (r *LinkRepository) Update(link *links.Link) error {
	rec := r.domainToRecord(link)
	res, err := r.db.NamedExec(`
        UPDATE links
		SET slug = :slug, target_url = :target_url, updated_at = :updated_at
        WHERE id = :id
    `, &rec)
	if err != nil {
		return fmt.Errorf("update link: %w", err)
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return links.ErrLinkNotFound
	}
	return nil
}

func (r *LinkRepository) Remove(link *links.Link) error {
	res, err := r.db.Exec("DELETE FROM links WHERE id = ?", link.ID())
	if err != nil {
		return fmt.Errorf("remove link: %w", err)
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return links.ErrLinkNotFound
	}
	return nil
}
