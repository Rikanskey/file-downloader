package repository

import (
	"file-downloader/internal/downloader_app"
	"strings"
)

type TaskStatus string

const (
	ready TaskStatus = "Ready"
	err   TaskStatus = "Error"
	wait  TaskStatus = "Wait"
)

type DownloadStatus struct {
	id         string
	taskStatus TaskStatus
	link       string
}

func (ds DownloadStatus) marshall() []string {
	return []string{ds.id, string(ds.taskStatus), ds.link}
}

func unmarshallTaskStatus(status string) downloader_app.TaskStatus {
	switch status {
	case string(downloader_app.Ready):
		return downloader_app.Ready
	case string(downloader_app.Wait):
		return downloader_app.Wait
	default:
		return downloader_app.Error
	}
}

type Task struct {
	Id    string
	Links []string
}

func (r Task) marshall() []string {
	return []string{r.Id, strings.Join(r.Links, ",")}
}
