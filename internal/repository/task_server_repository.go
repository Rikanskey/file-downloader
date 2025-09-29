package repository

import (
	"encoding/csv"
	"file-downloader/internal/app"
	"file-downloader/internal/config"
	"file-downloader/internal/downloader_app"
	"os"
	"strings"
	"sync"
)

type TaskServerRep struct {
	fileName string
}

func NewTaskServerRep(config config.TaskWriterConfig) *TaskServerRep {
	return &TaskServerRep{fileName: config.Path}
}

func (tw *TaskServerRep) WriteTask(request app.Task) {
	var mtx sync.Mutex
	mtx.Lock()

	file, err := os.OpenFile(tw.fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(Task{Id: request.Id, Links: request.Links}.marshall())
	if err != nil {
		return
	}

	mtx.Unlock()
}

func (tw *TaskServerRep) ReadAllTasks() ([]downloader_app.DownloadTask, error) {
	dts := make([]downloader_app.DownloadTask, 0)
	file, err := os.OpenFile(tw.fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return dts, err
	}
	defer file.Close()

	raws, err := csv.NewReader(file).ReadAll()
	for _, raw := range raws {
		dts = append(dts, downloader_app.DownloadTask{Id: raw[0], Links: strings.Split(raw[1], ",")})
	}

	return dts, nil
}
