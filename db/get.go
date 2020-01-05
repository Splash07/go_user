package db

import (
	"context"
	"time"

	"github.com/Splash07/go_user/model"
	"github.com/Splash07/go_user/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func LoginUser(req model.LoginReq) (*model.LoginResp, error) {
	var user model.User
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"email": req.Email, "password": utils.MD5(req.Password)}
	err := Client.Database(DbName).Collection(ColName).FindOne(ctx, filter).Decode(&user)
	token := utils.GenerateToken(user.ID, user.Phone, user.Email)
	resp := model.LoginResp{&user, token}
	return &resp, err
}

func FindUserByID(ID int) (*model.User, error) {
	var user model.User
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"id": ID}
	err := Client.Database(DbName).Collection(ColName).FindOne(ctx, filter).Decode(&user)

	return &user, err
}
