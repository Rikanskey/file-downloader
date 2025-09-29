package main

import (
	"file-downloader/internal/runner/server"
)

const configDir = "./config"

func main() {
	server.Start(configDir)
}
