package persistence

import (
	"ConnectionService/domain"
	"fmt"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func checkIfUserExist(userID string, transaction neo4j.Transaction) bool {
	result, _ := transaction.Run(
		"MATCH (existing_uer:USER) WHERE existing_uer.userID = $userID RETURN existing_uer.userID",
		map[string]interface{}{"userID": userID})

	if result != nil && result.Next() && result.Record().Values[0] == userID {
		return true
	}
	return false
}

func isUserPrivate(userID string, transaction neo4j.Transaction) (bool, error) {
	result, err := transaction.Run(
		"MATCH (existing_uer:USER) WHERE existing_uer.userID = $userID RETURN existing_uer.userID, existing_uer.isPrivate",
		map[string]interface{}{"userID": userID})

	if err != nil {
		return true, err
	}

	if result != nil && result.Next() {
		return result.Record().Values[1].(bool), nil
	}
	return true, err
}

func setUserPrivate(userID string, private bool, transaction neo4j.Transaction) (bool, error) {
	result, err := transaction.Run(
		"MATCH (u:USER) WHERE u.userID=$uID SET u.isPrivate=$private RETURN u.isPrivate ",
		map[string]interface{}{"uID": userID, "private": private})
	if err != nil {
		return false, err
	}
	if result != nil && result.Next() {
		return true, nil
	}
	return false, nil
}

func updateUser(user domain.UserConn, transaction neo4j.Transaction) (bool, error) {
	cypherText := "MATCH (u:USER) WHERE u.userID=$uID " +
		"SET u.isPrivate=$private " +
		"SET  u.username=$username " +
		"RETURN u.isPrivate, u.username "
	result, err := transaction.Run(
		cypherText,
		map[string]interface{}{"uID": user.UserID, "private": !user.IsPublic, "username": user.Username})
	if err != nil {
		return false, err
	}
	if result != nil && result.Next() && result.Record().Values[0].(bool) != user.IsPublic && result.Record().Values[1].(string) == user.Username {
		return true, nil
	}
	return false, nil
}

func checkIfFriendExist(userIDa, userIDb string, transaction neo4j.Transaction) bool {
	result, _ := transaction.Run(
		"MATCH (u1:USER) WHERE u1.userID=$uIDa "+
			"MATCH (u2:USER) WHERE u2.userID=$uIDb "+
			"MATCH (u1)-[r:FRIEND]->(u2) "+
			"RETURN r.date ",
		map[string]interface{}{"uIDa": userIDa, "uIDb": userIDb})

	if result != nil && result.Next() {
		return true
	}
	return false
}

func checkIfFriendRequestExist(userIDa, userIDb string, transaction neo4j.Transaction) bool {
	result, _ := transaction.Run(
		"MATCH (u1:USER) WHERE u1.userID=$uIDa "+
			"MATCH (u2:USER) WHERE u2.userID=$uIDb "+
			"MATCH (u1)-[r:REQUEST]->(u2) "+
			"RETURN r.date ",
		map[string]interface{}{"uIDa": userIDa, "uIDb": userIDb})

	if result != nil && result.Next() {
		return true
	}
	return false
}

func checkIfBlockExist(userIDa, userIDb string, transaction neo4j.Transaction) bool {
	result, _ := transaction.Run(
		"MATCH (u1:USER) WHERE u1.userID=$uIDa "+
			"MATCH (u2:USER) WHERE u2.userID=$uIDb "+
			"MATCH (u1)-[r:BLOCK]->(u2) "+
			"RETURN r.date ",
		map[string]interface{}{"uIDa": userIDa, "uIDb": userIDb})

	if result != nil && result.Next() {
		return true
	}
	return false
}

func blockUser(userIDa, userIDb string, transaction neo4j.Transaction) bool {

	dateNow := time.Now().Local().Unix()
	result, err := transaction.Run(
		"MATCH (u1:USER) WHERE u1.userID=$uIDa "+
			"MATCH (u2:USER) WHERE u2.userID=$uIDb "+
			"CREATE (u1)-[r:BLOCK {date: $dateNow}]->(u2) "+
			"RETURN r.date",
		map[string]interface{}{"uIDa": userIDa, "uIDb": userIDb, "dateNow": dateNow, "msgID": "newMsgID"})

	if err != nil {
		fmt.Println(err)
		return false
	}
	if result != nil && result.Next() {
		return true
	}
	return false
}

func removeFriend(userIDa, userIDb string, transaction neo4j.Transaction) bool {
	result, err := transaction.Run(
		"MATCH (u1:USER{userID: $uIDa})-[r:FRIEND]->(u2:USER{userID: $uIDb}) DELETE r RETURN u1.userID",
		map[string]interface{}{"uIDa": userIDa, "uIDb": userIDb})

	if err != nil {
		fmt.Println(err)
		return false
	}
	if result != nil && result.Next() {
		return true
	}
	return false
}

func addFriendRequest(userIDa, userIDb string, transaction neo4j.Transaction) error {
	_, err := transaction.Run(
		"MATCH (u1:USER) WHERE u1.userID=$uIDa "+
			"MATCH (u2:USER) WHERE u2.userID=$uIDb "+
			"CREATE (u1)-[r:REQUEST]->(u2) ",
		map[string]interface{}{"uIDa": userIDa, "uIDb": userIDb})
	return err
}

func removeFriendRequest(userIDa, userIDb string, transaction neo4j.Transaction) bool {
	result, err := transaction.Run(
		"MATCH (u1:USER{userID: $uIDa})-[r:REQUEST]->(u2:USER{userID: $uIDb}) DELETE r RETURN u1.userID",
		map[string]interface{}{"uIDa": userIDa, "uIDb": userIDb})

	if err != nil {
		fmt.Println(err)
		return false
	}
	if result != nil && result.Next() {
		return true
	}
	return false
}

func createFriendship(userIDa, userIDb string, transaction neo4j.Transaction) bool {
	_, err := transaction.Run(
		"MATCH (u1:USER) WHERE u1.userID=$uIDa "+
			"MATCH (u2:USER) WHERE u2.userID=$uIDb "+
			"CREATE (u1)-[r:FRIEND]->(u2)"+
			"CREATE (u2)-[r2:FRIEND]->(u1)",
		map[string]interface{}{"uIDa": userIDa, "uIDb": userIDb})
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func unblockUser(userIDa, userIDb string, transaction neo4j.Transaction) bool {
	result, err := transaction.Run(
		"MATCH (u1:USER{userID: $uIDa})-[r:BLOCK]->(u2:USER{userID: $uIDb}) DELETE r RETURN u1.userID",
		map[string]interface{}{"uIDa": userIDa, "uIDb": userIDb})

	if err != nil {
		fmt.Println(err)
		return false
	}
	if result != nil && result.Next() {
		return true
	}
	return false
}

func getFriendsOfFriends(userID string, transaction neo4j.Transaction) ([]*domain.UserRecommendation, error) {
	result, err := transaction.Run(
		"MATCH (u1:USER)-[:FRIEND]->(u2:USER)<-[:FRIEND]-(u3:USER) "+
			"MATCH (u1:USER)-[:FRIEND]->(u4:USER)<-[:FRIEND]-(u3:USER) "+
			"WHERE u1.userID=$uID AND u3.userID<>$uID AND u4.userID<>$uID AND u4.userID<>u3.userID "+
			"AND NOT exists((u1)-[:FRIEND]-(u3)) "+
			"AND NOT exists((u1)-[:BLOCK]-(u3)) "+
			"RETURN distinct u1.username as komePreporucujem ,u3.userID, u3.isPrivate, u3.username, count(u2) as mutual "+
			"ORDER BY mutual "+
			"LIMIT 20 ",
		map[string]interface{}{"uID": userID})

	if err != nil {
		return nil, err
	}
	var recommendation []*domain.UserRecommendation
	for result.Next() {
		recommendation = append(recommendation, &domain.UserRecommendation{
			UserID:   result.Record().Values[1].(string),
			Username: result.Record().Values[3].(string),
			IsPublic: !(result.Record().Values[2].(bool)),
			IsMutual: true,
			Mutual:   int(result.Record().Values[4].(int64)),
		})
	}
	return recommendation, nil
}

func getFriendRecommendation(userID string, transaction neo4j.Transaction) ([]*domain.UserRecommendation, error) {
	result, err := transaction.Run(
		"MATCH (u1:USER) "+
			"MATCH (u2:USER)-[r:FRIEND]->(:USER) "+
			"WHERE u1.userID=$uID AND u2.userID<>$uID "+
			"AND NOT exists((u1)-[:FRIEND]-(u2)) "+
			"AND NOT exists((u1)-[:BLOCK]-(u2)) "+
			"RETURN distinct u2.userID, u2.isPrivate, u2.username, COUNT(r) as num_of_friends "+
			"ORDER BY num_of_friends DESC "+
			"LIMIT 20 ",
		map[string]interface{}{"uID": userID})

	if err != nil {
		return nil, err
	}

	var recommendation []*domain.UserRecommendation
	for result.Next() {
		recommendation = append(recommendation, &domain.UserRecommendation{
			UserID:   result.Record().Values[0].(string),
			Username: result.Record().Values[2].(string),
			IsPublic: !(result.Record().Values[1].(bool)),
			IsMutual: false,
			Mutual:   int(result.Record().Values[3].(int64)),
		})
		//&domain.UserConn{UserID: result.Record().Values[0].(string), IsPublic: result.Record().Values[1].(bool)})
	}
	return recommendation, nil
}

func clearGraphDB(transaction neo4j.Transaction) error {
	_, err := transaction.Run(
		"MATCH (n) DETACH DELETE n",
		map[string]interface{}{})
	return err
}

func initGraphDB(transaction neo4j.Transaction) error {
	_, err := transaction.Run(
		"CREATE  (fedor:USER{userID: \"62c22d86cd5a19cfe6b7e6c7\", username: \"fedor\", isPrivate : false}),  "+
			"(andrea:USER{userID: \"62c22d86cd5a19cfe6b7e6c8\", username: \"andrea\", isPrivate : false}),  "+
			"(saska:USER{userID: \"62c22d86cd5a19cfe6b7e6c9\", username: \"saska\", isPrivate : false}),  "+
			"(igor:USER{userID: \"62c22d86cd5a19cfe6b7e6c0\", username: \"igor\", isPrivate : false}),      "+
			"(fedor) -[:FRIEND{msgID:\"12022d86cd5a19cfe6b7e6c0\"}]-> (andrea), "+
			" (andrea) <-[:FRIEND{msgID:\"23022d86cd5a19cfe6b7e6c0\"}]- (fedor),  "+
			"(fedor) -[:FRIEND{msgID:\"33022d86cd5a19cfe6b7e6c0\"}]-> (igor), "+
			" (igor) -[:FRIEND{msgID:\"43022d86cd5a19cfe6b7e6c0\"}]-> (fedor),   "+
			" (fedor) -[:BLOCK]-> (saska),  "+
			"(andrea) -[:BLOCK]-> (saska)  ",
		map[string]interface{}{})
	return err
}
