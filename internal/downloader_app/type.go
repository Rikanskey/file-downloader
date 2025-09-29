package downloader_app

type TaskStatus string

const (
	Ready TaskStatus = "Ready"
	Error TaskStatus = "Error"
	Wait  TaskStatus = "Wait"
)

type DownloadTask struct {
	Id    string
	Links []string
}

type DownloadStatusTask struct {
	Id     string
	Status TaskStatus
	Link   string
}
