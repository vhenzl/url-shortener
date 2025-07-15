package infrastructure

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
	"github.com/vhenzl/url-shortener/internal/domain/links"
	"github.com/vhenzl/url-shortener/internal/infrastructure/testutil"
)

type LinkRepositoryTestSuite struct {
	suite.Suite
	db         *sqlx.DB
	repository *LinkRepository
}

func TestLinkRepository(t *testing.T) {
	suite.Run(t, new(LinkRepositoryTestSuite))
}

func (s *LinkRepositoryTestSuite) SetupTest() {
	s.db = testutil.NewTestDB(s.T())
	err := testutil.RunMigrations(s.db.DB)
	s.Require().NoError(err)
	s.repository = NewLinkRepository(s.db)
}

func (s *LinkRepositoryTestSuite) TestAddAndGetByID() {
	// Create a test link
	id := links.LinkID(uuid.New())
	slug := links.Slug("test-slug")
	url := links.URL("https://example.com")
	now := time.Now().UTC()
	link := links.NewLink(id, slug, url, now, now)

	// Add to repository
	err := s.repository.Add(link)
	s.Require().NoError(err)

	// Get by ID
	retrieved, err := s.repository.GetByID(id)
	s.Require().NoError(err)
	s.Equal(link.ID(), retrieved.ID())
	s.Equal(link.Slug(), retrieved.Slug())
	s.Equal(link.TargetURL(), retrieved.TargetURL())
}

func (s *LinkRepositoryTestSuite) TestGetBySlug() {
	// Create a test link
	id := links.LinkID(uuid.New())
	slug := links.Slug("test-slug")
	url := links.URL("https://example.com")
	now := time.Now().UTC()
	link := links.NewLink(id, slug, url, now, now)

	// Add to repository
	err := s.repository.Add(link)
	s.Require().NoError(err)

	// Get by slug
	retrieved, err := s.repository.GetBySlug(slug)
	s.Require().NoError(err)
	s.Equal(link.ID(), retrieved.ID())
	s.Equal(link.Slug(), retrieved.Slug())
	s.Equal(link.TargetURL(), retrieved.TargetURL())
}

func (s *LinkRepositoryTestSuite) TestGetByID_NotFound() {
	id := links.LinkID(uuid.New())
	_, err := s.repository.GetByID(id)
	s.ErrorIs(err, links.ErrLinkNotFound)
}

func (s *LinkRepositoryTestSuite) TestGetBySlug_NotFound() {
	slug := links.Slug("non-existent")
	_, err := s.repository.GetBySlug(slug)
	s.ErrorIs(err, links.ErrLinkNotFound)
}
