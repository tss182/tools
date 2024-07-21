package tools

import (
	"context"
	"os"
	"strconv"
	"strings"
	"time"
)

func MongoDBCtx() (context.Context, context.CancelFunc) {
	timeStr := strings.TrimSpace(os.Getenv("MONGODB_TIMEOUT"))
	if timeStr == "" {
		timeStr = "5"
	}
	timeInt, _ := strconv.Atoi(timeStr)
	timeout := time.Duration(timeInt) * time.Minute
	return context.WithTimeout(context.Background(), timeout)
}
