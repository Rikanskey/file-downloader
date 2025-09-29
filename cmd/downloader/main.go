package main

import "file-downloader/internal/runner/downloader"

const configDir = "./config"

func main() {
	downloader.StartDownloader(configDir)
}
