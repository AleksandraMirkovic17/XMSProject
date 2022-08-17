package api

import (
	pb "github.com/dislinked/common/proto/job_service"
	"github.com/dislinked/job_service/domain"
)

func mapNewJob(pbJob *pb.Job) *domain.JobOffer {
	jobDomain := &domain.JobOffer{
		JobId:          pbJob.JobID,
		PublisherId:    pbJob.PublisherId,
		DatePosted:     pbJob.DatePosted,
		DateValid:      pbJob.DateValid,
		CompanyName:    pbJob.CompanyName,
		Position:       pbJob.Position,
		JobDescription: pbJob.JobDescription,
	}

	for _, skill := range pbJob.RequiredSkills {
		jobDomain.RequiredSkills = append(jobDomain.RequiredSkills, skill)

	}

	return jobDomain
}

func mapJobToPb(job *domain.JobOffer) *pb.Job {
	jobPB := &pb.Job{
		JobID:          job.JobId,
		PublisherId:    job.PublisherId,
		DatePosted:     job.DatePosted,
		DateValid:      job.DateValid,
		CompanyName:    job.CompanyName,
		Position:       job.Position,
		JobDescription: job.JobDescription,
	}

	for _, skill := range job.RequiredSkills {
		jobPB.RequiredSkills = append(jobPB.RequiredSkills, skill)
	}

	return jobPB

}
