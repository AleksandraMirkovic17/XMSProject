package persistence

import (
	"context"
	pb "github.com/dislinked/common/proto/job_service"
	"github.com/dislinked/job_service/domain"
)

type JobStore interface {
	Get(ctx context.Context, jobId string) (*domain.JobOffer, error)
	GetAll(ctx context.Context) ([]*domain.JobOffer, error)
	Insert(ctx context.Context, profile *domain.JobOffer) error
	Update(ctx context.Context, profile *domain.JobOffer) (bool, error)
	Search(ctx context.Context, search string) ([]*domain.JobOffer, error)
	GetUserJobOffers(ctx context.Context, userID string) ([]*domain.JobOffer, error)
	Delete(ctx context.Context, jobId string) (bool, error)
	Init()
	CreateUser(ctx context.Context, userID string) (*pb.ActionResult, error)
	UpdateUserSkills(ctx context.Context, userID string, skills []string) (*pb.ActionResult, error)
	GetRecommendationJobOffer(ctx context.Context, userID string) ([]*domain.JobOffer, error)
}
