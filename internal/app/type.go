package app

import "encoding/json"

type TaskStatus string

const (
	Ready TaskStatus = "Ready"
	Error TaskStatus = "Error"
	Wait  TaskStatus = "Wait"
)

type DownloadStatus struct {
	Id         string
	TaskStatus TaskStatus
	Link       string
}

type Task struct {
	Id    string
	Links []string
}

func (r Task) Marshall() ([]byte, error) {
	return json.Marshal(r)
}
