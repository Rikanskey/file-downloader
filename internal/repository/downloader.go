package repository

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type FileSaver interface {
	Save(data FileData) error
}

type UrlDownloader struct {
	fileSaver FileSaver
}

type fileWriter struct {
	dirName string
}

type FileData struct {
	filename string
	body     []byte
}

func NewUrlDownloader(dirName string) *UrlDownloader {
	return &UrlDownloader{fileSaver: fileWriter{dirName}}
}

func (fw fileWriter) Save(data FileData) error {
	file, err := os.Create(fw.dirName + data.filename)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()
	_, err = file.Write(data.body)
	return err
}

func (ud *UrlDownloader) httpDownload(url string) ([]byte, error) {
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)

	if err != nil {
	}
	return body, nil
}

func (ud *UrlDownloader) Download(url string) error {
	body, err := ud.httpDownload(url)
	if err != nil {
		return err
	}

	s := strings.Split(url, "/")
	fl := FileData{filename: s[len(s)-1], body: body}

	err = ud.fileSaver.Save(fl)

	return err
}
