package persistence

import (
	"../model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateTask(task *model.Task) error {
	_, err := collection.InsertOne(ctx, task)
	return err
}

func GetAll() ([]*model.Task, error) {
	filter := bson.D{{}}
	return filterTasks(filter)
}

func filterTasks(filter interface{}) ([]*model.Task, error) {
	var tasks []*model.Task
	curr, err := collection.Find(ctx, filter)
	if err != nil {
		return tasks, err
	}

	for curr.Next(ctx) {
		var t model.Task
		err := curr.Decode(&t)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, &t)
	}
	if err := curr.Err(); err != nil {
		return tasks, err
	}
	_ = curr.Close(ctx)
	if len(tasks) == 0 {
		return tasks, mongo.ErrNoDocuments
	}
	return tasks, nil
}
