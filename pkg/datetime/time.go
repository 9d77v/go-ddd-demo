package datetime

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func FormatTimeStamp(timestamp *timestamppb.Timestamp) string {
	return timestamp.AsTime().Local().Format(time.RFC3339)
}
