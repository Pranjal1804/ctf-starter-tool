package models

import (
	"context"
	"errors"
	"time"

	"ctf-toolkit-backend/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// User represents a user in the system.
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username  string             `bson:"username" json:"username"`
	Password  string             `bson:"password" json:"-"`
	Email     string             `bson:"email" json:"email"`
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at" json:"updated_at"`
}

// SaveUser saves a user to the database
func SaveUser(user User) error {
	collection := database.GetDatabase().Collection("users")

	// Set timestamps
	now := primitive.NewDateTimeFromTime(time.Now())
	user.CreatedAt = now
	user.UpdatedAt = now

	// Check if username already exists
	var existingUser User
	err := collection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&existingUser)
	if err == nil {
		return errors.New("username already exists")
	}

	// Check if email already exists
	err = collection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		return errors.New("email already exists")
	}

	// Insert the user
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}

// FindUserByUsername finds a user by username
func FindUserByUsername(username string) (User, error) {
	collection := database.GetDatabase().Collection("users")

	var user User
	err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, errors.New("user not found")
		}
		return user, err
	}

	return user, nil
}

// FindUserByEmail finds a user by email
func FindUserByEmail(email string) (User, error) {
	collection := database.GetDatabase().Collection("users")

	var user User
	err := collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, errors.New("user not found")
		}
		return user, err
	}

	return user, nil
}

// UpdateUser updates a user in the database
func UpdateUser(user User) error {
	collection := database.GetDatabase().Collection("users")

	// Set update timestamp
	user.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}

// DeleteUser deletes a user from the database
func DeleteUser(userID primitive.ObjectID) error {
	collection := database.GetDatabase().Collection("users")

	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": userID})
	return err
}