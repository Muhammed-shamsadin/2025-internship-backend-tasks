package data

import (
	"2025-internship-backend-tasks/Task-Management-API/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) CreateUser(user *models.User) error {
	_, err := UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return errors.New("failed to create user")
	}

	return nil
}


func (us *UserService) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := UserCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}


func (us *UserService) UsernameExists(username string) (bool, error) {
	count, err := UserCollection.CountDocuments(context.TODO(), bson.M{"username": username})
	if err != nil {
		return false, errors.New("failed to check if username exists")
	}
	return count > 0, nil
}