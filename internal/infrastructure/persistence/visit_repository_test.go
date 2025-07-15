package infrastructure

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
	"github.com/vhenzl/url-shortener/internal/domain/links"
	"github.com/vhenzl/url-shortener/internal/domain/visits"
	"github.com/vhenzl/url-shortener/internal/infrastructure/testutil"
)

type VisitRepositoryTestSuite struct {
	suite.Suite
	db              *sqlx.DB
	visitRepository *VisitRepository
	linkRepository  *LinkRepository
}

func TestVisitRepository(t *testing.T) {
	suite.Run(t, new(VisitRepositoryTestSuite))
}

func (s *VisitRepositoryTestSuite) SetupTest() {
	s.db = testutil.NewTestDB(s.T())
	err := testutil.RunMigrations(s.db.DB)
	s.Require().NoError(err)
	s.visitRepository = NewVisitRepository(s.db)
	s.linkRepository = NewLinkRepository(s.db)
}

func (s *VisitRepositoryTestSuite) createTestLink() *links.Link {
	id := links.LinkID(uuid.New())
	slug := links.Slug("test-slug")
	url := links.URL("https://example.com")
	now := time.Now().UTC()
	link := links.NewLink(id, slug, url, now, now)
	err := s.linkRepository.Add(link)
	s.Require().NoError(err)
	return link
}

func (s *VisitRepositoryTestSuite) TestAddAndGetByID() {
	// First create a link
	link := s.createTestLink()

	// Create a test visit
	id := visits.VisitID(uuid.New())
	now := time.Now().UTC()
	visit := visits.NewVisit(id, link.ID(), now)

	// Add to repository
	err := s.visitRepository.Add(visit)
	s.Require().NoError(err)

	// Get by ID
	retrieved, err := s.visitRepository.GetByID(id)
	s.Require().NoError(err)
	s.Equal(visit.ID(), retrieved.ID())
	s.Equal(visit.LinkID(), retrieved.LinkID())
	s.Equal(visit.VisitedAt().Unix(), retrieved.VisitedAt().Unix())
}

func (s *VisitRepositoryTestSuite) TestGetAllByLinkID() {
	// Create a link
	link := s.createTestLink()

	// Create multiple visits
	allVisits := make([]*visits.Visit, 3)
	for i := range allVisits {
		id := visits.NewVisitID()
		now := time.Now().Add(time.Duration(i) * time.Hour).UTC() // Different times
		allVisits[i] = visits.NewVisit(id, link.ID(), now)
		err := s.visitRepository.Add(allVisits[i])
		s.Require().NoError(err)
	}

	// Get all visits for the link
	retrieved, err := s.visitRepository.GetAllByLinkID(link.ID())
	s.Require().NoError(err)
	s.Len(retrieved, len(allVisits))

	// Verify visits are returned in descending order by visited_at
	for i := 1; i < len(retrieved); i++ {
		s.True(retrieved[i-1].VisitedAt().After(retrieved[i].VisitedAt()))
	}
}

func (s *VisitRepositoryTestSuite) TestGetAllByLinkID_NoVisits() {
	// Create a link
	link := s.createTestLink()

	// Get visits for link with no visits
	retrieved, err := s.visitRepository.GetAllByLinkID(link.ID())
	s.Require().NoError(err)
	s.Empty(retrieved)
}

func (s *VisitRepositoryTestSuite) TestGetByID_NotFound() {
	id := visits.VisitID(uuid.New())
	_, err := s.visitRepository.GetByID(id)
	s.ErrorIs(err, visits.ErrVisitNotFound)
}
