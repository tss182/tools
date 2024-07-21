package tools

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CountPipeline(ctx context.Context, col *mongo.Collection, pipeline []interface{}) (int64, error) {
	pipeline = append(pipeline, bson.D{{OperatorCount, "count"}})
	var countResult []CountPage
	cur, err := col.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, err
	}
	for cur.Next(ctx) {
		var doc CountPage
		err = cur.Decode(&doc)
		if err != nil {
			return 0, err
		}
		countResult = append(countResult, doc)
	}
	var count int64
	if len(countResult) > 0 {
		count = countResult[0].Count
	}
	return count, nil
}
