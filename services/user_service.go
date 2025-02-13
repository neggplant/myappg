package services

import (
	"context"
	"myappg/models"
	"myappg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	collection *mongo.Collection
}

func NewUserService() *UserService {
	db := utils.GetDB()
	return &UserService{
		collection: db.Collection("users"),
	}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User

	cursor, err := s.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) CreateUser(user *models.User) error {
	_, err := s.collection.InsertOne(context.Background(), user)
	return err
}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.Logger.Error(err.Error())
		return nil, err
	}
	// utils.Logger.Error("error1")
	var user models.User
	err = s.collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserService) UpdateUser(id string, user *models.User) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"name":  user.Name,
			"email": user.Email,
		},
	}

	_, err = s.collection.UpdateOne(context.Background(), bson.M{"_id": objectID}, update)
	return err
}

func (s *UserService) DeleteUser(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = s.collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	return err
}
