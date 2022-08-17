package domain

import (
	"context"
	pb "github.com/dislinked/common/proto/job_service"
)

type JobStore interface {
	Get(ctx context.Context, jobId string) (*JobOffer, error)
	GetAll(ctx context.Context) ([]*JobOffer, error)
	Insert(ctx context.Context, profile *JobOffer) error
	Update(ctx context.Context, profile *JobOffer) (bool, error)
	Search(ctx context.Context, search string) ([]*JobOffer, error)
	GetUserJobOffers(ctx context.Context, userID string) ([]*JobOffer, error)
	Delete(ctx context.Context, jobId string) (bool, error)
	Init()
	CreateUser(ctx context.Context, userID string, username string) (*pb.ActionResult, error)
	DeleteUser(ctx context.Context, userID string) (*pb.ActionResult, error)
	UpdateUserSkills(ctx context.Context, userID string, skills []string) (*pb.ActionResult, error)
	GetRecommendationJobOffer(ctx context.Context, userID string) ([]*JobOffer, error)
}
