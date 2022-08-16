package persistence

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

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

func checkIfUserExists(userID string, transaction neo4j.Transaction) bool {
	cypherText := string(database) + "MATCH (existing_uer:USER) WHERE existing_uer.userID = $userID RETURN existing_uer.userID"
	result, _ := transaction.Run(
		cypherText,
		map[string]interface{}{"userID": userID})

	if result != nil && result.Next() && result.Record().Values[0] == userID {
		return true
	}
	return false
}
