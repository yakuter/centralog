package entry

import (
	"time"
)

type Entry struct {
	ID           uint      `gorm:"primarykey"`
	Project      string    `json:"project"`
	Hostname     string    `json:"hostname"`
	OS           string    `json:"os"`
	Platform     string    `json:"platform"`
	Architecture string    `json:"architecture"`
	PipelineName string    `json:"pipeline_name"`
	PipelineID   string    `json:"pipeline_id"`
	Level        string    `json:"level"`
	Message      string    `json:"message"`
	Timestamp    time.Time `json:"timestamp"`
}
