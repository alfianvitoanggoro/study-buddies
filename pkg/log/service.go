package log

import (
	"context"

	"github.com/AlfianVitoAnggoro/study-buddies/pkg/elasticsearch"
)

const (
	INDEX_LOG_ERROR    = "log_error"
	INDEX_LOG_ACTIVITY = "log_activity"
	INDEX_LOG_LOGIN    = "log_login"
)

func InsertErrorLog(ctx context.Context, log *LogError) error {
	// return nil
	return elasticsearch.Insert(ctx, INDEX_LOG_ERROR, log)
}

func InsertActivityLog(ctx context.Context, log *LogActivity) error {
	// return nil
	return elasticsearch.Insert(ctx, INDEX_LOG_ACTIVITY, log)
}

func InsertLoginLog(ctx context.Context, log *LogLogin) error {
	// return nil
	return elasticsearch.Insert(ctx, INDEX_LOG_LOGIN, log)
}
