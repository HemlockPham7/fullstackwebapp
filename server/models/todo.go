package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

type TodoRepository interface {
	getTodos(ctx context.Context) ([]*Todo, error)
	getTodo(ctx context.Context, objectID primitive.ObjectID) (*Todo, error)
	createTodo(ctx context.Context, todo *Todo) (*Todo, error)
	updateTodo(ctx context.Context, objectID primitive.ObjectID) (*Todo, error)
	deleteTodo(ctx context.Context, objectID primitive.ObjectID) error
}
