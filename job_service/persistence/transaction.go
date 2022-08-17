package persistence

import (
	"github.com/dislinked/job_service/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

const database = "USE JOBS "

func createUserNode() {

}

func initGraphDB(transaction neo4j.Transaction) error {
	_, err := transaction.Run(
		"CREATE DATABSE jobs IF NOT EXISTS ",
		map[string]interface{}{})
	if err != nil {
		cypherText := string(database) + " CREATE (Advokat:JOB{position:advokat})"
		_, err = transaction.Run(
			cypherText,
			map[string]interface{}{})
	}

	return err
}

func deleteNode(userID string, transaction neo4j.Transaction) error {
	cypherText := string(database) + "MATCH(existing_uer:USERJOB) WHERE existing_uer.userID = $userID DETACH DELETE existing_uer"
	_, err := transaction.Run(
		cypherText,
		map[string]interface{}{"userID": userID})
	return err
}

func checkIfUserExists(userID string, transaction neo4j.Transaction) bool {
	cypherText := string(database) + "MATCH (existing_uer:USERJOB) WHERE existing_uer.userID = $userID RETURN existing_uer.userID"
	result, _ := transaction.Run(
		cypherText,
		map[string]interface{}{"userID": userID})

	if result != nil && result.Next() && result.Record().Values[0] == userID {
		return true
	}
	return false
}

func checkIfSkillExists(skill string, transaction neo4j.Transaction) bool {
	cypherText := string(database) + "MATCH (existing_skill:SKILL) WHERE existing_skill.name=$skill RETURN existing_skill"
	result, _ := transaction.Run(
		cypherText,
		map[string]interface{}{"skill": skill})

	if result != nil && result.Next() && result.Record().Values[0] == skill {
		return true
	}

	return false

}

func checkIfJobExists(jobId string, transaction neo4j.Transaction) bool {
	cypherText := string(database) + "MATCH (existing_job:JOB) WHERE existing_job.jobID=$jobID RETURN existing_job"
	result, _ := transaction.Run(
		cypherText,
		map[string]interface{}{"jobID": jobId})

	if result != nil && result.Next() && result.Record().Values[0] == jobId {
		return true
	}

	return false
}

func createSkillNode(skill string, transaction neo4j.Transaction) bool {
	if !checkIfSkillExists(skill, transaction) {
		cypherText := string(database) + "CREATE (:SKILL{name:$skill})"
		_, err := transaction.Run(
			cypherText,
			map[string]interface{}{"skill": skill},
		)
		if err != nil {
			return false
		}
		return true
	} else {
		return false
	}

}

func createJobNode(job *domain.JobOffer, transaction neo4j.Transaction) error {
	cypherText := string(database) + "CREATE (:JOB{jobID:$jobID, publisherID:$publisherID, datePosted:$datePosted, dateValid:$dateValid, companyName:$companyName, position:$position, description:$description })"
	_, err := transaction.Run(
		cypherText,
		map[string]interface{}{"jobID": job.JobId, "publisherID": job.PublisherId, "datePosted": job.DatePosted, "dateValid": job.DateValid, "companyName": job.CompanyName, "position": job.Position, "description": job.JobDescription},
	)
	return err

}

func connectSkillAndJob(jobID string, skill string, transaction neo4j.Transaction) (bool, error) {
	println("skill connectiong", skill)
	cypherText := string(database) + "MATCH (j:JOB) WHERE j.jobID=$jobID " +
		"MATCH (s:SKILL) WHERE s.name=$skill " +
		"CREATE (j) -[:REQUIRE]-> (s) " +
		"CREATE (s) -[:isREQUIRED] -> (j) " +
		"RETURN j.jobID, s.name"
	result, err := transaction.Run(
		cypherText,
		map[string]interface{}{"jobID": jobID, "skill": skill})
	if err != nil {
		return false, err
	}

	if result != nil && result.Next() && result.Record().Values[0].(string) == jobID && result.Record().Values[1].(string) == skill {
		return true, nil
	}
	return false, nil

}
