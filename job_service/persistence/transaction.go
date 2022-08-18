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
