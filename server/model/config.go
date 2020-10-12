package model

//Global variable for accessing mongo db configs
var (
	Config dbConfig
)

//database configuration settings
type dbConfig struct {
	MongoURL      string
	MongoPort     string
	MongoDBName   string
	MongoUsername string
	MongoPass     string
	HostName      string
}

//GetConfigFilePath : to load mongodb configurations
func GetConfigFilePath() string {
	return "config/config.toml"
}
