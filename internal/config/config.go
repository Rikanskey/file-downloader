package config

import "github.com/spf13/viper"

type (
	ServerConfig struct {
		Http                HttpConfig
		RequestLogger       TaskWriterConfig
		DownloadTasksConfig DownloadTasksConfig
	}

	DownloaderConfig struct {
		TaskWriter     TaskWriterConfig
		DownloadTasks  DownloadTasksConfig
		FilesDirectory FilesDirectoryConfig
	}

	HttpConfig struct {
		Host string
		Port string
	}

	TaskWriterConfig struct {
		Path string
	}
	DownloadTasksConfig struct {
		Path string
	}
	FilesDirectoryConfig struct {
		Path string
	}
)

const (
	downloader = "application-downloader"
	server     = "application-server"
)

func NewServerConfig(configDir string) (*ServerConfig, error) {
	var config ServerConfig

	if err := parseConfigFile(configDir, server); err != nil {
		return nil, err
	}

	if err := unmarshalServer(&config.Http); err != nil {
		return nil, err
	}

	if err := unmarshalTaskWriter(&config.RequestLogger); err != nil {
		return nil, err
	}

	if err := unmarshalDownloadTasksConfig(&config.DownloadTasksConfig); err != nil {
		return nil, err
	}

	return &config, nil
}

func NewDownloaderConfig(configDir string) (*DownloaderConfig, error) {
	var config DownloaderConfig

	if err := parseConfigFile(configDir, downloader); err != nil {
		return nil, err
	}

	if err := unmarshalDownloaderConfig(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func unmarshalTaskWriter(config *TaskWriterConfig) error {
	return viper.UnmarshalKey("request-logger", &config)
}

func unmarshalDownloadTasksConfig(config *DownloadTasksConfig) error {
	return viper.UnmarshalKey("download-tasks", &config)
}

func parseConfigFile(configDir string, configName string) error {
	viper.AddConfigPath(configDir)
	viper.SetConfigName(configName)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.SetConfigName(configName)
	return viper.MergeInConfig()
}

func unmarshalServer(config *HttpConfig) error {
	return viper.UnmarshalKey("http", &config)
}

func unmarshalDownloaderConfig(config *DownloaderConfig) error {
	if err := viper.UnmarshalKey("files-directory", &config.FilesDirectory); err != nil {
		return err
	}

	if err := unmarshalDownloadTasksConfig(&config.DownloadTasks); err != nil {
		return err
	}

	return unmarshalTaskWriter(&config.TaskWriter)
}
