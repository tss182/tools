package tools

import "time"

type (
	HistoryRequest struct {
		ID        string                 `bson:"_id"`
		Title     string                 `bson:"title"`
		Method    string                 `bson:"method"`
		Url       string                 `bson:"url"`
		Header    map[string]interface{} `bson:"header"`
		Request   interface{}            `bson:"request"`
		Response  interface{}            `bson:"response"`
		Duration  float64                `bson:"duration"`
		CreatedAt time.Time              `bson:"created_at"`
	}

	CountPage struct {
		Count int64 `bson:"count"`
	}
)

const (
	ShouldTypeQuery = "query"
	ShouldTypeJson  = "json"
	ShouldTypeForm  = "form"

	OperatorCount = "$count"
)
