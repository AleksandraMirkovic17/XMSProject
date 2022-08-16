package persistence

import (
	"context"
	"fmt"
	pb "github.com/dislinked/common/proto/job_service"
	"github.com/dislinked/job_service/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type JobDBStore struct {
	jobDB *neo4j.Driver
}

func NewJobDBStore(client *neo4j.Driver) domain.JobStore {
	return &JobDBStore{
		jobDB: client,
	}
}

func (j JobDBStore) Get(ctx context.Context, jobId string) (*domain.JobOffer, error) {
	//TODO implement me
	panic("implement me")
}

func (j JobDBStore) GetAll(ctx context.Context) ([]*domain.JobOffer, error) {
	//TODO implement me
	panic("implement me")
}

func (j JobDBStore) Insert(ctx context.Context, profile *domain.JobOffer) error {
	//TODO implement me
	panic("implement me")
}

func (j JobDBStore) Update(ctx context.Context, profile *domain.JobOffer) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (j JobDBStore) Search(ctx context.Context, search string) ([]*domain.JobOffer, error) {
	//TODO implement me
	panic("implement me")
}

func (j JobDBStore) GetUserJobOffers(ctx context.Context, userID string) ([]*domain.JobOffer, error) {
	//TODO implement me
	panic("implement me")
}

func (j JobDBStore) Delete(ctx context.Context, jobId string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (j *JobDBStore) Init() {
	session := (*j.jobDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		errInit := initGraphDB(transaction)
		return nil, errInit
	})

	if err != nil {
		fmt.Println("Connection Graph Database INIT - Failed", err.Error())
	} else {
		fmt.Println("Connection Graph Database INIT - Successfully")
	}
}

func (j JobDBStore) CreateUser(ctx context.Context, userID string, username string) (*pb.ActionResult, error) {
	println("Creatig user job node!")
	session := (*j.jobDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	result, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		actionResult := &pb.ActionResult{}
		if checkIfUserExists(userID, transaction) {
			actionResult.Status = 406
			actionResult.Msg = "error user with ID:" + userID + " already exist"
			println("Already existing user!")
			return actionResult, nil
		}

		_, err := transaction.Run(
			database+"CREATE (new_user:USERJOB{userID:$userID, username:$username})",
			map[string]interface{}{"userID": userID, "username": username})

		if err != nil {
			println("error while creating new node with ID:" + userID)
			actionResult.Msg = "error while creating new node with ID:" + userID
			actionResult.Status = 501
			return actionResult, err
		}
		actionResult.Msg = "Successfully created new node with ID:" + userID
		println("Successfully created new node with ID:" + userID)
		actionResult.Status = 201

		return actionResult, err

	})
	if result == nil {
		return &pb.ActionResult{Msg: "error", Status: 500}, err
	} else {
		return result.(*pb.ActionResult), err
	}
}

func (j JobDBStore) UpdateUserSkills(ctx context.Context, userID string, skills []string) (*pb.ActionResult, error) {
	//TODO implement me
	panic("implement me")
}

func (j JobDBStore) GetRecommendationJobOffer(ctx context.Context, userID string) ([]*domain.JobOffer, error) {
	//TODO implement me
	panic("implement me")
}
