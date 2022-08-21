package persistence

import (
	"context"
	"fmt"
	pb "github.com/dislinked/common/proto/job_service"
	"github.com/dislinked/job_service/domain"
	"github.com/dislinked/job_service/infrastructure/errors"
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
	session := (*j.jobDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	jobById, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		if !(checkIfJobExists(jobId, transaction)) {
			return nil, errors.NewCustomError("Job with id " + jobId + " does not exist!")
		}
		job, err1 := getJobById(jobId, transaction)
		if err1 != nil {
			return nil, err1
		}
		return job, nil
	},
	)
	if err != nil {
		return nil, err
	}
	return jobById.(*domain.JobOffer), nil
}
func (j JobDBStore) GetAll(ctx context.Context) ([]*domain.JobOffer, error) {
	session := (*j.jobDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	allJobs, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		job, err1 := getAllJobs(transaction)
		if err1 != nil {
			return nil, err1
		}
		return job, nil
	},
	)
	if err != nil {
		return nil, err
	}
	return allJobs.([]*domain.JobOffer), nil

}

func (j JobDBStore) Insert(ctx context.Context, job *domain.JobOffer) error {
	session := (*j.jobDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		if !checkIfJobExists(job.JobId, transaction) {
			errJob := createJobNode(job, transaction)
			if errJob != nil {
				return nil, errJob
			}
			for _, skill := range job.RequiredSkills {
				createSkillNode(skill, transaction)
				connectSkillAndJob(job.JobId, skill, transaction)

			}

		}
		return nil, nil
	})
	return err
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
	session := (*j.jobDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	allJobs, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		job, err1 := getJobsByPublisherId(userID, transaction)
		if err1 != nil {
			return nil, err1
		}
		return job, nil
	},
	)
	if err != nil {
		return nil, err
	}
	return allJobs.([]*domain.JobOffer), nil
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

func (j JobDBStore) DeleteUser(ctx context.Context, userID string) (*pb.ActionResult, error) {
	println("Deleting user job node!")
	session := (*j.jobDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	result, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		actionResult := &pb.ActionResult{}
		if !checkIfUserExists(userID, transaction) {
			actionResult.Status = 406
			actionResult.Msg = "error user with ID:" + userID + " does not exist"
			println("User does not exist!")
			return actionResult, nil
		}
		err := deleteNode(userID, transaction)

		if err != nil {
			println("error while deleting node with ID:" + userID)
			actionResult.Msg = "error while deleting node with ID:" + userID
			actionResult.Status = 501
			return actionResult, err
		}
		actionResult.Msg = "Successfully deleted node with ID:" + userID
		println("Successfully deleted node with ID:" + userID)
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

func (j JobDBStore) UpdateUser(ctx context.Context, node domain.UserJobNode) (*pb.ActionResult, error) {
	session := (*j.jobDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		if !checkIfUserExists(node.UserID, transaction) {
			return &pb.ActionResult{
				Status: 400,
				Msg:    "User does not exsist!",
			}, nil
		}
		isUpdated, err := updateUser(node, transaction)
		if err != nil || !isUpdated {
			return &pb.ActionResult{
				Status: 400,
				Msg:    "It is not updated!",
			}, err
		}
		return &pb.ActionResult{
			Status: 200,
			Msg:    "Successfully update userjob node!",
		}, nil
	})

	if err != nil {
		return &pb.ActionResult{
			Status: 400,
			Msg:    "It is not updated! " + err.Error(),
		}, err

	}

	return &pb.ActionResult{
		Status: 200,
		Msg:    "Successfully update userjob node!",
	}, nil

}

func (j JobDBStore) GetRecommendationJobOffer(ctx context.Context, userID string) ([]*domain.JobOffer, error) {
	//TODO implement me
	panic("implement me")
}
