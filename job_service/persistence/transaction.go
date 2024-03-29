package persistence

import (
	"github.com/dislinked/job_service/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

var database = "USE JOBS "

func createUserNode() {

}

func initGraphDB(transaction neo4j.Transaction) error {
	_, err := transaction.Run(
		"CREATE DATABASE jobs IF NOT EXISTS ",
		map[string]interface{}{})
	if err != nil {
		println("Database uspesno kreiran")
		cypherText := string(database) + " CREATE (Advokat:JOB{position:advokat})"
		_, err = transaction.Run(
			cypherText,
			map[string]interface{}{})
	}
	println("database mozda nije uspesno kreiran")
	if err != nil {
		println("Moramo koristiti deaultni database")
		database = ""
	}

	return err
}

func getJobById(jobID string, transaction neo4j.Transaction) (*domain.JobOffer, error) {
	cypherText := string(database) + "MATCH (existing_job:JOB) " +
		"WHERE existing_job.jobID=$jobID " +
		"RETURN existing_job.jobID, existing_job.publisherID, existing_job.companyName, existing_job.datePosted, existing_job.dateValid, existing_job.description,  existing_job.position"
	result, err := transaction.Run(
		cypherText,
		map[string]interface{}{"jobID": jobID})
	if err != nil || result == nil {
		return nil, err
	}
	result.Next()
	foundJob := &domain.JobOffer{
		JobId:          result.Record().Values[0].(string),
		PublisherId:    result.Record().Values[1].(string),
		RequiredSkills: []string{},
		DatePosted:     result.Record().Values[3].(string),
		DateValid:      result.Record().Values[4].(string),
		CompanyName:    result.Record().Values[2].(string),
		Position:       result.Record().Values[6].(string),
		JobDescription: result.Record().Values[5].(string),
	}

	//get skills
	cypherTextSkills := string(database) + "MATCH (job:JOB)-[r:REQUIRE]->(s:SKILL) " +
		"WHERE job.jobID=$jobID " +
		"RETURN s.name "
	resultSkills, err := transaction.Run(
		cypherTextSkills,
		map[string]interface{}{"jobID": jobID})
	if err != nil {
		return nil, err
	}
	for resultSkills.Next() {
		println("Ubacivanje skilla u job: " + resultSkills.Record().Values[0].(string))
		foundJob.RequiredSkills = append(foundJob.RequiredSkills, resultSkills.Record().Values[0].(string))
	}
	return foundJob, nil

}

func getAllJobs(transaction neo4j.Transaction) ([]*domain.JobOffer, error) {
	cypherText := string(database) + "MATCH (j:JOB) RETURN j.jobID"
	result, err := transaction.Run(
		cypherText,
		map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	var jobs []*domain.JobOffer

	for result.Next() {
		job, err := getJobById(result.Record().Values[0].(string), transaction)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func getJobsByPublisherId(publisherId string, transaction neo4j.Transaction) ([]*domain.JobOffer, error) {
	cypherText := string(database) + "MATCH (j:JOB) WHERE j.publisherID=$publisherID RETURN j.jobID "
	result, err := transaction.Run(
		cypherText,
		map[string]interface{}{"publisherID": publisherId})
	if err != nil {
		return nil, err
	}
	var jobs []*domain.JobOffer

	for result.Next() {
		job, err := getJobById(result.Record().Values[0].(string), transaction)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	return jobs, nil

}

func getRecommendedJobs(userID string, transaction neo4j.Transaction) ([]*domain.JobOffer, error) {
	cypherText := string(database) + "MATCH (u:USERJOB)-[k:KNOWS]->(s:SKILL)-[r:isREQUIRED]->(j:JOB) " +
		"WHERE u.userID=$userID " +
		"RETURN DISTINCT j.jobID, count(s) as numberOfSkills " +
		"ORDER BY numberOfSkills DESC"
	result, err := transaction.Run(
		cypherText,
		map[string]interface{}{"userID": userID})
	if err != nil {
		return nil, err
	}
	var jobs []*domain.JobOffer

	for result.Next() {
		job, err := getJobById(result.Record().Values[0].(string), transaction)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	return jobs, nil

}

func getJobsBySearch(params string, transaction neo4j.Transaction) ([]*domain.JobOffer, error) {
	cypherText := string(database) + "MATCH (j:JOB) " +
		"WHERE j.position CONTAINS $search " +
		"OR j.companyName CONTAINS $search " +
		"RETURN j.jobID "
	result, err := transaction.Run(
		cypherText,
		map[string]interface{}{"search": params})
	if err != nil {
		return nil, err
	}
	var jobs []*domain.JobOffer

	for result.Next() {
		job, err := getJobById(result.Record().Values[0].(string), transaction)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func deleteNode(userID string, transaction neo4j.Transaction) error {
	cypherText := string(database) + "MATCH(existing_uer:USERJOB) WHERE existing_uer.userID = $userID DETACH DELETE existing_uer"
	_, err := transaction.Run(
		cypherText,
		map[string]interface{}{"userID": userID})
	return err
}

func checkIfUserExists(userID string, transaction neo4j.Transaction) bool {
	cypherText := string(database) + "MATCH (existing_uer:USERJOB) " +
		"WHERE existing_uer.userID = $userID RETURN existing_uer.userID"
	result, _ := transaction.Run(
		cypherText,
		map[string]interface{}{"userID": userID})

	if result != nil && result.Next() && result.Record().Values[0] == userID {
		return true
	}
	return false
}

func checkIfSkillExists(skill string, transaction neo4j.Transaction) bool {
	cypherText := string(database) + "MATCH (existing_skill:SKILL) WHERE existing_skill.name=$skill RETURN existing_skill.name"
	result, _ := transaction.Run(
		cypherText,
		map[string]interface{}{"skill": skill})

	if result != nil && result.Next() && result.Record().Values[0] == skill {
		return true
	}

	return false

}

func checkIfJobExists(jobId string, transaction neo4j.Transaction) bool {
	cypherText := string(database) + "MATCH (existing_job:JOB) WHERE existing_job.jobID=$jobID RETURN existing_job.jobID"
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
		println("skill: ", skill, " does not exist! creating")
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

func updateUser(node domain.UserJobNode, transaction neo4j.Transaction) (bool, error) {
	println("updating username")
	cypherText := string(database) +
		"MATCH (existing_uer:USERJOB) " +
		"WHERE existing_uer.userID = $userID " +
		"SET existing_uer.username = $username " +
		"RETURN existing_uer.username"
	result, err := transaction.Run(
		cypherText,
		map[string]interface{}{"userID": node.UserID, "username": node.Username})
	if err != nil || result == nil {
		return false, err
	}
	if !result.Next() || result.Record().Values[0].(string) != node.Username {
		println("Not updated username!")
		return false, err
	}
	println("getting old skills")
	oldSkills, err := getUserSkills(node.UserID, transaction)
	if err != nil {
		println("Couldn't get old skills for user!")
		return false, err
	}
	//delete old skills
	for _, oldSkill := range oldSkills {
		if !checkIfSkillPresent(oldSkill, node.Skills) {
			println("old skills: ", oldSkill, " is not present in new deleting")
			isRemoved, err := removeUserSkill(node.UserID, oldSkill, transaction)
			if err != nil || !isRemoved {
				return false, err
			}
			println("Removed skill " + oldSkill)
		}
	}

	//add new skills
	for _, newSkill := range node.Skills {
		if !checkIfSkillPresent(newSkill, oldSkills) {
			createSkillNode(newSkill, transaction)
			println("New skill does not existed before in skills")
			addedNew, err := addSkillToUser(node.UserID, newSkill, transaction)
			if err != nil || !addedNew {
				return false, err
			}
			println("Added new skill " + newSkill)
		}

	}

	return true, nil

}

func addSkillToUser(userID string, newSkill string, transaction neo4j.Transaction) (bool, error) {
	cypherText := string(database) + " MATCH (u:USERJOB) WHERE u.userID=$userID " +
		"MATCH (s:SKILL) WHERE s.name=$name " +
		"CREATE (u)-[:KNOWS] ->(s) " +
		"CREATE (s)-[:isKNOWN] ->(u) " +
		"RETURN u.userID, s.name "
	result, err := transaction.Run(
		cypherText,
		map[string]interface{}{"userID": userID, "name": newSkill})
	if err != nil {
		return false, err
	}
	if result != nil && result.Next() && result.Record().Values[0].(string) == userID && result.Record().Values[1].(string) == newSkill {
		return true, nil
	}
	return false, nil
}

func removeUserSkill(userID string, oldSkill string, transaction neo4j.Transaction) (bool, error) {
	cypherText := string(database) + " MATCH (u:USERJOB)-[k:KNOWS]->(s:SKILL) " +
		"MATCH (s:SKILL)-[k2:isKNOWN]->(u:USERJOB) " +
		"WHERE u.userID=$userID AND s.name=$skill " +
		"DELETE k " +
		"DELETE k2 " +
		"RETURN s.name "
	result, err := transaction.Run(
		cypherText,
		map[string]interface{}{"userID": userID, "skill": oldSkill})
	if err != nil {
		return false, err
	}
	if result != nil && result.Next() && result.Record().Values[0].(string) == oldSkill {
		return true, nil
	}
	return false, nil

}

func checkIfSkillPresent(oldSkill string, newSkills []string) bool {
	for _, skill := range newSkills {
		if oldSkill == skill {
			return true
		}
	}
	return false
}

func getUserSkills(userID string, transaction neo4j.Transaction) ([]string, error) {
	cypherText := string(database) + "MATCH (u:USERJOB) - [k:KNOWS]->(s:SKILL) " +
		"WHERE u.userID=$userID " +
		"return s.name"
	result, err := transaction.Run(
		cypherText,
		map[string]interface{}{"userID": userID})

	var userSkills []string

	if err != nil || result == nil {
		return nil, err
	}

	for result.Next() {
		userSkills = append(userSkills, result.Record().Values[0].(string))
		print(",old skill: " + result.Record().Values[0].(string))
	}
	println(" , len: ", len(userSkills))
	return userSkills, nil
}
