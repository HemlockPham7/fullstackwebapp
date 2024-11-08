package repositories

import (
	"context"
	"log"

	"github.com/HemlockPham7/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepository struct {
	db *mongo.Collection
}

func NewTodoRepository(db *mongo.Collection) models.TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (r *TodoRepository) GetTodos(ctx context.Context) ([]*models.Todo, error) {
	todos := []*models.Todo{}

	cursor, err := r.db.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatalf("Find error")
		return nil, err
	}
	defer cursor.Close(context.Background()) // Ensure the cursor is closed when done

	for cursor.Next(context.Background()) {
		var todo *models.Todo
		if err := cursor.Decode(&todo); err != nil {
			log.Fatalf("Decode error")
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *TodoRepository) GetTodo(ctx context.Context, objectID primitive.ObjectID) (*models.Todo, error) {
	filter := bson.M{"_id": objectID}
	todo := &models.Todo{}

	err := r.db.FindOne(context.Background(), filter).Decode(todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *TodoRepository) CreateTodo(ctx context.Context, todo *models.Todo) (*models.Todo, error) {
	insertResult, err := r.db.InsertOne(context.Background(), todo)
	if err != nil {
		return nil, err
	}

	todo.ID = insertResult.InsertedID.(primitive.ObjectID)
	return todo, nil
}

func (r *TodoRepository) UpdateTodo(ctx context.Context, objectID primitive.ObjectID) (*models.Todo, error) {
	filter := bson.M{"_id": objectID}

	update := bson.M{"$set": bson.M{"completed": true}}
	_, err := r.db.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	todo, err := r.GetTodo(ctx, objectID)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *TodoRepository) DeleteTodo(ctx context.Context, objectID primitive.ObjectID) error {
	filter := bson.M{"_id": objectID}

	_, err := r.db.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}
