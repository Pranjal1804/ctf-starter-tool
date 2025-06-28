package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Log represents the structure of a log entry in MongoDB
type Log struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id,omitempty"`
	ToolName  string             `bson:"tool_name"`
	Timestamp int64              `bson:"timestamp"`
	Input     string             `bson:"input"`
	Output    string             `bson:"output"`
}