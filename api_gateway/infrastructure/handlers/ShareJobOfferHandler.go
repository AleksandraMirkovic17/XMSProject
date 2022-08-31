package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dislinked/api_gateway/domain"
	handler "github.com/dislinked/api_gateway/infrastructure/api"
	clients "github.com/dislinked/api_gateway/infrastructure/services"
	"github.com/dislinked/api_gateway/startup/config"
	pbJob "github.com/dislinked/common/proto/job_service"
	pb "github.com/dislinked/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type ShareJobOfferHandler struct {
	userServiceAddress string
	jobServiceAddress  string
	config             *config.Config
}

func (f ShareJobOfferHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/share/job", f.ShareJobOffer)
	if err != nil {
		println(err.Error())
		panic(err)
	}
}

func NewShareJobOfferHandler(c *config.Config) handler.Handler {
	return &ShareJobOfferHandler{
		userServiceAddress: fmt.Sprintf("%s:%s", c.UserHost, c.UserPort),
		jobServiceAddress:  fmt.Sprintf("%s:%s", c.JobHost, c.JobPort),
		config:             c,
	}
}

func (f ShareJobOfferHandler) ShareJobOffer(w http.ResponseWriter, r *http.Request, _ map[string]string) {
	var shareJobRequest domain.ShareJobRequest
	err := json.NewDecoder(r.Body).Decode(&shareJobRequest)
	if err != nil {
		http.Error(w, "Error decoding share job offer request! "+err.Error(), http.StatusBadRequest)
		return
	}
	userService := clients.NewUserClient(f.userServiceAddress)
	jobService := clients.NewJobClient(f.jobServiceAddress)

	respus, err := userService.ShareJobOffer(context.TODO(), &pb.ShareJobRequest{ShareJob: &pb.ShareJob{
		ApiToken: shareJobRequest.ApiToken,
		Job: &pb.Job{
			JobID:          "",
			PublisherId:    "",
			RequiredSkills: shareJobRequest.Job.RequiredSkills,
			DatePosted:     shareJobRequest.Job.DatePosted,
			DateValid:      shareJobRequest.Job.DateValid,
			CompanyName:    shareJobRequest.Job.CompanyName,
			Position:       shareJobRequest.Job.Position,
			JobDescription: shareJobRequest.Job.JobDescription,
		},
	}})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		println("Error occured")
		return
	}
	if !respus.Valid {
		w.WriteHeader(http.StatusForbidden)
		println("It is not valid")
		return
	}

	respjob, error := jobService.CreateJob(context.TODO(), &pbJob.CreateJobRequest{
		Job: &pbJob.Job{
			JobID:          "",
			PublisherId:    respus.Job.PublisherId,
			RequiredSkills: respus.Job.RequiredSkills,
			DatePosted:     respus.Job.DatePosted,
			DateValid:      respus.Job.DateValid,
			CompanyName:    respus.Job.CompanyName,
			Position:       respus.Job.Position,
			JobDescription: respus.Job.JobDescription,
		}})

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		println("Job service error occured")
		return
	}

	jobCreated := domain.Job{
		JobId:          respjob.Job.JobID,
		PublisherId:    respjob.Job.PublisherId,
		RequiredSkills: respjob.Job.RequiredSkills,
		DatePosted:     respjob.Job.DatePosted,
		DateValid:      respjob.Job.DateValid,
		CompanyName:    respjob.Job.CompanyName,
		Position:       respjob.Job.Position,
		JobDescription: respjob.Job.JobDescription,
	}

	jobCreatedJson, _ := json.Marshal(jobCreated)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jobCreatedJson)
	if err != nil {
		println("writing response error:", err.Error())
		w.WriteHeader(http.StatusConflict)
		return
	}

}
