package api

import (
	pb "github.com/dislinked/common/proto/job_service"
	"github.com/dislinked/job_service/application"
	"golang.org/x/net/context"
)

type JobHandler struct {
	pb.UnimplementedJobServiceServer
	service *application.JobService
}

func NewConnectionHandler(service *application.JobService) *JobHandler {
	return &JobHandler{
		UnimplementedJobServiceServer: pb.UnimplementedJobServiceServer{},
		service:                       service,
	}
}

func (handler *JobHandler) CreateJob(ctx context.Context, request *pb.CreateJobRequest) (*pb.CreateJobResponse, error) {
	job := mapNewJob((*request).Job)

	newJob, err := handler.service.Insert(ctx, job)

	if err != nil {
		return &pb.CreateJobResponse{Job: nil, Message: "Failed to create job node!"}, err
	}
	return &pb.CreateJobResponse{
		Job:     mapJobToPb(newJob),
		Message: "Successfully created!",
	}, nil

}
