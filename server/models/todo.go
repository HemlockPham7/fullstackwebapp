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
	GetTodos(ctx context.Context) ([]*Todo, error)
	GetTodo(ctx context.Context, objectID primitive.ObjectID) (*Todo, error)
	CreateTodo(ctx context.Context, todo *Todo) (*Todo, error)
	UpdateTodo(ctx context.Context, objectID primitive.ObjectID) (*Todo, error)
	DeleteTodo(ctx context.Context, objectID primitive.ObjectID) error
}
