package dbactionsservice

import (
	"context"
	"fmt"

	dbtoolkit "ApiExtention.com/src/services/db_service/funcs/toolkit"
	req_res_types "ApiExtention.com/src/types"
	"go.mongodb.org/mongo-driver/bson"
)

// Query token in DB
func Query_Token(tokenToSearch string) bool {
	context_custom := context.Background()

	client, err := dbtoolkit.Gen_ServerAPIClient(&context_custom)
	if err != nil {
		fmt.Println(err)
	}

	coll := client.Database("MinerAds").Collection("users")
	filter := bson.D{{Key: "token", Value: tokenToSearch}}

	result := req_res_types.User{}
	err = coll.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		fmt.Println(err)
		return false
	}

	defer dbtoolkit.Close_DataBase_Connection(client, &context_custom)
	return true
}

// Insert a New User in DB
func Isert_New_User_in_DB(newUser *req_res_types.User) error {
	context_custom := context.Background()

	client, errGen := dbtoolkit.Gen_ServerAPIClient(&context_custom)
	if errGen != nil {
		fmt.Println(errGen)
		return errGen
	}

	coll := client.Database("MinerAds").Collection("users")

	new_user := req_res_types.User{
		Data_User:       newUser.Data_User,
		Token:           newUser.Token,
		Subscription_ID: newUser.Data_User.Subscription_ID,
	}

	_, errInsert := coll.InsertOne(context_custom, new_user)
	if errInsert != nil {
		fmt.Println("Error em inserir o usuário!: " + errInsert.Error())
		return errInsert
	}

	defer dbtoolkit.Close_DataBase_Connection(client, &context_custom)
	return nil
}

// Delete an User DB
func Delete_User(subscription *req_res_types.KiwifyResponse) error {
	context_custom := context.Background()

	client, errGen := dbtoolkit.Gen_ServerAPIClient(&context_custom)
	if errGen != nil {
		fmt.Println(errGen)
		return errGen
	}

	coll := client.Database("MinerAds").Collection("users")
	to_delete := bson.D{{Key: "subscription_id", Value: subscription.Subscription_ID}}

	_, errInsert := coll.DeleteMany(context_custom, to_delete)
	if errInsert != nil {
		fmt.Println("Error ao excluir um usuário!: " + errInsert.Error())
		return errInsert
	}

	defer dbtoolkit.Close_DataBase_Connection(client, &context_custom)
	return nil
}

// Delete Token for overdue subscriptions
func Clean_Token_User(subscription *req_res_types.KiwifyResponse) error {
	context_custom := context.Background()

	client, errGen := dbtoolkit.Gen_ServerAPIClient(&context_custom)
	if errGen != nil {
		fmt.Println(errGen)
		return errGen
	}

	coll := client.Database("MinerAds").Collection("users")
	filter := bson.D{{Key: "subscription_id", Value: subscription.Subscription_ID}}
	update_to := bson.D{{Key: "$set", Value: bson.D{{Key: "subscription_id", Value: ""}}}}

	_, errInsert := coll.UpdateOne(context_custom, filter, update_to)
	if errInsert != nil {
		fmt.Println("Error ao excluir um usuário!: " + errInsert.Error())
		return errInsert
	}

	defer dbtoolkit.Close_DataBase_Connection(client, &context_custom)
	return nil
}
