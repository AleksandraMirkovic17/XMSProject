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
func (handler *JobHandler) GetJobOffer(ctx context.Context, request *pb.GetJobRequest) (*pb.GetJobResponse, error) {
	job, err := handler.service.GetJobOffer(ctx, request.JobID)
	if err != nil {
		return nil, err
	}
	return &pb.GetJobResponse{Job: mapJobToPb(job)}, nil
}

func (handler *JobHandler) GetAllJobs(ctx context.Context, request *pb.EmptyJobRequest) (*pb.GetAllJobsResponse, error) {
	jobs, err := handler.service.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllJobsResponse{Jobs: []*pb.Job{}}
	for _, job := range jobs {
		response.Jobs = append(response.Jobs, mapJobToPb(job))
	}
	return response, nil
}

func (handler *JobHandler) GetUserJobs(ctx context.Context, request *pb.CreateUserRequest) (*pb.GetAllJobsResponse, error) {
	jobs, err := handler.service.GetAllByPublisher(ctx, request.UserID)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllJobsResponse{Jobs: []*pb.Job{}}
	for _, job := range jobs {
		response.Jobs = append(response.Jobs, mapJobToPb(job))
	}
	return response, nil
}

//rpc SearchJob(SearchJobRequest) returns(GetAllJobsResponse){
func (handler *JobHandler) SearchJob(ctx context.Context, request *pb.SearchJobRequest) (*pb.GetAllJobsResponse, error) {
	jobs, err := handler.service.GetAllBySearchParams(ctx, request.Param)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllJobsResponse{Jobs: []*pb.Job{}}
	for _, job := range jobs {
		response.Jobs = append(response.Jobs, mapJobToPb(job))
	}
	return response, nil
}

func (handler *JobHandler) GetRecommendationJob(ctx context.Context, request *pb.GetRecommendedJobsRequest) (*pb.GetAllJobsResponse, error) {
	jobs, err := handler.service.GetRecommendedJobs(ctx, request.UserID)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllJobsResponse{Jobs: []*pb.Job{}}
	for _, job := range jobs {
		response.Jobs = append(response.Jobs, mapJobToPb(job))
	}
	return response, nil
}
