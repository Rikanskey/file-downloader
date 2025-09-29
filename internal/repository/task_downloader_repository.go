package repository

import (
	"context"
	"encoding/csv"
	"file-downloader/internal/app"
	"file-downloader/internal/config"
	"file-downloader/internal/downloader_app"
	"go/types"
	"log"
	"os"
	"sync"
)

type TaskDownloaderRep struct {
	fileName string
}

func NewTaskReader(cfg config.DownloadTasksConfig) *TaskDownloaderRep {
	return &TaskDownloaderRep{fileName: cfg.Path}
}

func (tdr *TaskDownloaderRep) ReadAllTasksId() ([]string, error) {
	keys := make(map[string]bool)
	result := make([]string, 0)
	file, err := os.OpenFile(tdr.fileName, os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Println(err)
		return result, err
	}
	defer file.Close()

	raws, err := csv.NewReader(file).ReadAll()

	if err != nil {
		log.Println(err)
		return result, err
	}

	for _, raw := range raws {
		if !keys[raw[0]] {
			keys[raw[0]] = true
			result = append(result, raw[0])
		}
	}

	return result, nil
}

func (tdr *TaskDownloaderRep) CreateTask(task downloader_app.DownloadTask) error {
	var mtx sync.Mutex
	mtx.Lock()
	file, err := os.OpenFile(tdr.fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Println(err)
		return err
	}
	raws, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Println(err)
		return err
	}
	file.Close()
	file, err = os.OpenFile(tdr.fileName, os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Println(err)
		return err
	}

	for _, link := range task.Links {
		raws = append(raws, DownloadStatus{id: task.Id, taskStatus: wait, link: link}.marshall())
	}
	err = csv.NewWriter(file).WriteAll(raws)
	mtx.Unlock()
	return err
}

func (tdr *TaskDownloaderRep) ReadTask(ctx context.Context, taskId string) ([]app.DownloadStatus, error) {
	result := make([]app.DownloadStatus, 0)
	file, err := os.OpenFile(tdr.fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return result, err
	}

	defer file.Close()

	raws, err := csv.NewReader(file).ReadAll()
	for _, raw := range raws {
		if raw[0] == taskId {
			result = append(result, app.DownloadStatus{Id: taskId, TaskStatus: app.TaskStatus(raw[1]), Link: raw[2]})
		}
	}
	return result, nil
}

func (tdr *TaskDownloaderRep) ReadTaskByStatus(status downloader_app.TaskStatus) ([]downloader_app.DownloadStatusTask, error) {
	result := make([]downloader_app.DownloadStatusTask, 0)
	file, err := os.OpenFile(tdr.fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return result, err
	}
	defer file.Close()
	raws, err := csv.NewReader(file).ReadAll()
	for _, raw := range raws {
		if raw[1] == string(status) {
			result = append(result, downloader_app.DownloadStatusTask{Id: raw[0],
				Status: unmarshallTaskStatus(raw[1]), Link: raw[2]})
		}
	}

	return result, nil
}

func (tdr *TaskDownloaderRep) UpdateStatusByLink(link string, status downloader_app.TaskStatus) error {
	var mtx sync.Mutex
	mtx.Lock()
	defer mtx.Unlock()
	file, err := os.OpenFile(tdr.fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Println(err)
		return err
	}

	raws, err := csv.NewReader(file).ReadAll()
	file.Close()
	file, err = os.OpenFile(tdr.fileName, os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	for _, raw := range raws {
		if raw[2] == link {
			raw[1] = string(status)
			err = writer.WriteAll(raws)
			if err != nil {
				log.Println(err)
				return err
			}
			return nil
		}
	}

	return types.Error{Msg: "Not found by link"}
}
