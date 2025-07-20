package data

import (
	"2025-internship-backend-tasks/Task-Management-API/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var tasks = []models.Task{}

type TaskService struct{}

func NewTaskService() *TaskService {
	return &TaskService{}
}

// GetAllTasks
func (ts *TaskService) GetAllTasks() []models.Task {
	var tasks []models.Task
	cursor, err := TaskCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return tasks
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var task models.Task
		if err := cursor.Decode(&task); err == nil {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

// GetTaskByID returns a task by its ID
func (ts *TaskService) GetTaskByID(id string) (*models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid task ID")
	}
	var task models.Task
	err = TaskCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		return nil, errors.New("task not found")
	}
	return &task, nil
}

// Update tasks
func (ts *TaskService) UpdateTask(id string, updatedTask models.Task) (*models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid task ID")
	}
	update := bson.M{"$set": bson.M{}}
	if updatedTask.Title != "" {
		update["$set"].(bson.M)["title"] = updatedTask.Title
	}
	if updatedTask.Description != "" {
		update["$set"].(bson.M)["description"] = updatedTask.Description
	}
	if updatedTask.Status != "" {
		update["$set"].(bson.M)["status"] = updatedTask.Status
	}
	if !updatedTask.DueDate.IsZero() {
		update["$set"].(bson.M)["due_date"] = updatedTask.DueDate
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	res := TaskCollection.FindOneAndUpdate(context.TODO(), bson.M{"_id": objID}, update, opts)
	var task models.Task
	err = res.Decode(&task)
	if err != nil {
		return nil, errors.New("task not found or update failed")
	}
	return &task, nil
}

func (ts *TaskService) AddTask(task models.Task) (*models.Task, error) {
	// Generate a new ObjectID if not provided
	if task.ID == "" {
		taskObjID := primitive.NewObjectID()
		task.ID = taskObjID.Hex()
	}
	objID, err := primitive.ObjectIDFromHex(task.ID)
	if err != nil {
		return nil, errors.New("invalid task ID")
	}
	doc := bson.M{
		"_id":         objID,
		"title":       task.Title,
		"description": task.Description,
		"due_date":    task.DueDate,
		"status":      task.Status,
	}
	_, err = TaskCollection.InsertOne(context.TODO(), doc)
	if err != nil {
		return nil, errors.New("failed to add task")
	}
	return &task, nil
}

func (ts *TaskService) DeleteTask(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task ID")
	}
	res, err := TaskCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return errors.New("failed to delete task")
	}
	if res.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}
