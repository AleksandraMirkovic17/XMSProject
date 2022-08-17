package application

import (
	"context"
	"fmt"
	pb "github.com/dislinked/common/proto/job_service"
	userService "github.com/dislinked/common/proto/user_service"
	"github.com/dislinked/job_service/domain"
	"github.com/dislinked/job_service/startup/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobService struct {
	store      domain.JobStore
	UserClient userService.UserServiceClient
}

func NewJobService(store domain.JobStore, c *config.Config) *JobService {
	return &JobService{
		store:      store,
		UserClient: NewUserClient(fmt.Sprintf("%s:%s", c.UserHost, c.UserPort)),
	}

}

func (service *JobService) CreateUser(ctx context.Context, userID string, username string) (*pb.ActionResult, error) {
	return service.store.CreateUser(ctx, userID, username)
}

func (service *JobService) DeleteUser(ctx context.Context, userID string) (*pb.ActionResult, error) {
	return service.store.DeleteUser(ctx, userID)
}

func (service *JobService) Insert(ctx context.Context, offer *domain.JobOffer) (*domain.JobOffer, error) {
	if (*offer).JobId == "" {
		//TODO: nekako ubaciti id?
		id := primitive.NewObjectID()
		(*offer).JobId = id.Hex()
	}
	err := service.store.Insert(ctx, offer)
	if err != nil {
		return nil, err
	}
	return offer, err

}
