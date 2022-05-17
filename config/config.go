package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	ENVIRONMENT          string
	MONGO_DB_NAME        string
	MONGO_USERNAME       string
	MONGO_PASSWORD       string
	MONGO_URL            string
	PORT                 int
	IAMPORT_API_KEY      string
	IAMPORT_SECRET_KEY   string
	IAMPORT_IMP_CODE     string
	PUSH_SERVER_URL      string
	PUSH_NAVIGATE_URL    string
	POSTGRES_URL         string
	POSTGRES_USER        string
	POSTGRES_PASSWORD    string
	POSTGRES_DB_NAME     string
	GRPC_PORT            int
	AMPLITUDE_API_KEY    string
	REDIS_URL            string
	ELASTICSEARCH_URL    string
	ELASTICSEARCH_APIKEY string
	OMNIOUS_KEY          string
}

func GetConfiguration(env string) Configuration {
	configuration := Configuration{}
	err := gonfig.GetConf(getFileName(env), &configuration)
	if err != nil {
		fmt.Println(err)
		os.Exit(500)
	}

	return configuration
}

func getFileName(env string) string {
	filename := []string{"/", "config.", env, ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))

	return filePath
}
