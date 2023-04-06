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
	Pipeline     string    `json:"pipeline"`
	Level        string    `json:"level"`
	Message      string    `json:"message"`
	Timestamp    time.Time `json:"timestamp"`
}
