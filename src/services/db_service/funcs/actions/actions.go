package dbactionsservice

import (
	"context"
	"fmt"

	dbtoolkit "ApiExtention.com/src/services/db_service/funcs/toolkit"
	reqkiwifywhtype "ApiExtention.com/src/types/req_kiwify_wh"
	usertype "ApiExtention.com/src/types/user"
	"go.mongodb.org/mongo-driver/bson"
)

// Query token in DB
func Query_Token(tokenToSearch string) bool {
	context_custom := context.Background()

	client, err := dbtoolkit.Gen_ServerAPIClient(&context_custom)
	if err != nil {
		fmt.Println(err)
	}
	defer dbtoolkit.Close_DataBase_Connection(client, &context_custom)

	coll := client.Database("MinerAds").Collection("users")
	filter := bson.D{{Key: "token", Value: tokenToSearch}}

	result := usertype.User{}

	if err = coll.FindOne(context.Background(), filter).Decode(&result); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// Insert a New User in DB
func Isert_New_User_in_DB(newUser *usertype.User) error {
	context_custom := context.Background()

	client, errGen := dbtoolkit.Gen_ServerAPIClient(&context_custom)
	if errGen != nil {
		fmt.Println(errGen)
		return errGen
	}

	coll := client.Database("MinerAds").Collection("users")

	new_user := usertype.User{
		Data_User:       newUser.Data_User,
		Token:           newUser.Token,
		Subscription_ID: newUser.Data_User.Subscription_id,
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
func Delete_User(subscription *reqkiwifywhtype.Req_Kiwify_Wh_Type) error {
	context_custom := context.Background()

	client, errGen := dbtoolkit.Gen_ServerAPIClient(&context_custom)
	if errGen != nil {
		fmt.Println(errGen)
		return errGen
	}

	coll := client.Database("MinerAds").Collection("users")
	to_delete := bson.D{{Key: "subscription_id", Value: subscription.Subscription_id}}

	_, errInsert := coll.DeleteMany(context_custom, to_delete)
	if errInsert != nil {
		fmt.Println("Error ao excluir um usuário!: " + errInsert.Error())
		return errInsert
	}

	defer dbtoolkit.Close_DataBase_Connection(client, &context_custom)
	return nil
}

// Delete Token for overdue subscriptions
func Clean_Token_User(subscription *reqkiwifywhtype.Req_Kiwify_Wh_Type) error {
	context_custom := context.Background()

	client, errGen := dbtoolkit.Gen_ServerAPIClient(&context_custom)
	if errGen != nil {
		fmt.Println(errGen)
		return errGen
	}

	coll := client.Database("MinerAds").Collection("users")
	filter := bson.D{{Key: "subscription_id", Value: subscription.Subscription_id}}
	update_to := bson.D{{Key: "$set", Value: bson.D{{Key: "token", Value: ""}}}}

	_, errInsert := coll.UpdateOne(context_custom, filter, update_to)
	if errInsert != nil {
		fmt.Println("Error ao excluir um usuário!: " + errInsert.Error())
		return errInsert
	}

	defer dbtoolkit.Close_DataBase_Connection(client, &context_custom)
	return nil
}

// Update Token for renewed subscriptions
func Update_Token_User(user *usertype.User) error {
	context_custom := context.Background()

	client, errGen := dbtoolkit.Gen_ServerAPIClient(&context_custom)
	if errGen != nil {
		fmt.Println(errGen)
		return errGen
	}

	coll := client.Database("MinerAds").Collection("users")
	filter := bson.D{{Key: "subscription_id", Value: user.Subscription_ID}}
	update_to := bson.D{{Key: "$set", Value: bson.D{{Key: "token", Value: user.Token}}}}

	_, errInsert := coll.UpdateOne(context_custom, filter, update_to)
	if errInsert != nil {
		fmt.Println("Error ao excluir um usuário!: " + errInsert.Error())
		return errInsert
	}

	defer dbtoolkit.Close_DataBase_Connection(client, &context_custom)
	return nil
}
