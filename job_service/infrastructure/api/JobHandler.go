package api

import (
	pb "github.com/dislinked/common/proto/job_service"
	"github.com/dislinked/job_service/application"
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
