package repositories

import (
	"2025-internship-backend-tasks/Task-Management-API/domain/task"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTaskRepository struct {
	collection *mongo.Collection
}

func NewMongoTaskRepository(collection *mongo.Collection) *MongoTaskRepository {
	return &MongoTaskRepository{collection: collection}
}

func (r *MongoTaskRepository) Create(ctx context.Context, task *task.Task) error {
	_, err := r.collection.InsertOne(ctx, task)
	return err
}

func (r *MongoTaskRepository) FindByID(ctx context.Context, id string) (*task.Task, error) {
	var t task.Task
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *MongoTaskRepository) FindAll(ctx context.Context) ([]task.Task, error) {
	var tasks []task.Task
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var t task.Task
		if err := cursor.Decode(&t); err == nil {
			tasks = append(tasks, t)
		}
	}
	return tasks, nil
}

func (r *MongoTaskRepository) Update(ctx context.Context, id string, t *task.Task) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": t})
	return err
}

func (r *MongoTaskRepository) Delete(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
