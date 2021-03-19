package persistence

import (
	"time"

	"../model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func Completed(id string) error {
	docID, _ := primitive.ObjectIDFromHex(id)
	_, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": docID},
		bson.M{
			"$set": bson.M{"completed": true, "updated_at": time.Now()},
		},
	)
	return err

}

func ExistID(id string) error {
	docID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": docID}
	t := model.Task{}
	return collection.FindOne(ctx, filter).Decode(&t)

}

func Delete (id string) error {
	docID, _ := primitive.ObjectIDFromHex(id)
	_,err := collection.DeleteOne(
		ctx,
		bson.M{"_id": docID},
		)
		return err 
}
