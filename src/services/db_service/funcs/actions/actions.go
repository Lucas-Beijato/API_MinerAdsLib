package dbactionsservice

import (
	"context"
	"fmt"

	dbtoolkit "ApiExtention.com/src/services/db_service/funcs/toolkit"
	req_res_types "ApiExtention.com/src/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Token    string             `bson:"token"`
	isActive bool               `bson:"isActive"`
}

// Query token in DB
func Query_Token(tokenToSearch string) (bool, string) {
	context_custom := context.Background()

	client, err := dbtoolkit.Gen_ServerAPIClient(&context_custom)
	if err != nil {
		fmt.Println(err)
	}

	coll := client.Database("MinerAds").Collection("users")
	filter := bson.D{{Key: "token", Value: tokenToSearch}}

	result := User{}
	err = coll.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		fmt.Println(err)
		return false, "Token is not valid."
	}

	if !result.isActive {
		return true, "User is not active."
	}

	defer dbtoolkit.Close_DataBase_Connection(client, &context_custom)
	return true, "Valid User."
}

// Insert a New User in DB
func Isert_New_User_in_DB(newUser *req_res_types.InsertNewUserDB) error {
	context_custom := context.Background()

	client, errGen := dbtoolkit.Gen_ServerAPIClient(&context_custom)
	if errGen != nil {
		fmt.Println(errGen)
		return errGen
	}

	coll := client.Database("MinerAds").Collection("users")

	_, errInsert := coll.InsertOne(context_custom, newUser)
	if errInsert != nil {
		fmt.Println("Error em inserir o usuário!: " + errInsert.Error())
		return errInsert
	}

	defer dbtoolkit.Close_DataBase_Connection(client, &context_custom)
	return nil
}
