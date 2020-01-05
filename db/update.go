package db

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Splash07/go_user/model"
	"github.com/Splash07/go_user/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func AddUser(user *model.User) (interface{}, error) {
	user.Password = utils.MD5(user.Password)
	collection := Client.Database(DbName).Collection(ColName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, user)
	return res, err
}

func UpdateUser(userReq *model.UserUpdateReq) (interface{}, error) {
	var user map[string]interface{}

	bs, _ := json.Marshal(userReq)
	json.Unmarshal(bs, &user)
	if _, ok := user["password"]; ok {
		user["password"] = utils.MD5(user["password"].(string))
	}

	collection := Client.Database(DbName).Collection(ColName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"id": userReq.ID}
	update := bson.M{"$set": user}
	res, err := collection.UpdateOne(ctx, filter, update)
	return res, err
}
