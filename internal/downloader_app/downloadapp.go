package downloader_app

import (
	"log"
	"os"
	"slices"
	"time"
)

type TaskServerReader interface {
	ReadAllTasks() ([]DownloadTask, error)
}

type TaskDownloader interface {
	ReadTaskByStatus(status TaskStatus) ([]DownloadStatusTask, error)
	UpdateStatusByLink(link string, status TaskStatus) error
	ReadAllTasksId() ([]string, error)
	CreateTask(downloadTask DownloadTask) error
}

type DownloadApplication struct {
	fileStat         string
	interval         time.Duration
	taskServerReader TaskServerReader
	taskDownloader   TaskDownloader
	app              Application
}

func NewDownloadApplication(fileStat string, interval time.Duration, taskServerReader TaskServerReader,
	taskDownloader TaskDownloader, app Application) *DownloadApplication {
	return &DownloadApplication{
		fileStat:         fileStat,
		interval:         interval,
		taskServerReader: taskServerReader,
		taskDownloader:   taskDownloader,
		app:              app,
	}
}

func (da *DownloadApplication) downloadUrl(dst DownloadStatusTask) {
	err := da.app.Commands.CreateDownloadFileCommand.Handle(CreateDownloadFileCommand{Url: dst.Link})
	if err != nil {
		log.Println(err)
		err = da.taskDownloader.UpdateStatusByLink(dst.Link, Error)
	} else {
		err = da.taskDownloader.UpdateStatusByLink(dst.Link, Ready)
	}
	if err != nil {
		log.Println(err)
	}
}

func (da *DownloadApplication) download() error {
	dsts, err := da.taskDownloader.ReadTaskByStatus(Wait)
	if err != nil {
		return err
	}

	for _, dst := range dsts {
		go da.downloadUrl(dst)
	}

	return nil
}

func (da *DownloadApplication) checkActualTasks() {
	tasksFromServer, _ := da.taskServerReader.ReadAllTasks()
	tasksIDs, err := da.taskDownloader.ReadAllTasksId()
	for _, task := range tasksFromServer {
		if !slices.Contains(tasksIDs, task.Id) {
			err = da.taskDownloader.CreateTask(task)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (da *DownloadApplication) Start() error {
	da.checkActualTasks()
	err := da.download()
	if err != nil {
		log.Println(err)
		return err
	}
	initialStat, err := os.Stat(da.fileStat)
	if err != nil {
		log.Panicln(err)
		return err
	}
	for {
		stat, err := os.Stat(da.fileStat)
		if err != nil {
			log.Println(err)
			return err
		}
		if stat.Size() != initialStat.Size() {
			da.checkActualTasks()
			err = da.download()
			if err != nil {
				log.Println(err)
				return err
			}
			initialStat = stat
		}
		time.Sleep(da.interval * time.Second)
	}
}
