package config

var (
	Currency Config
)

type Config struct {
	Debug        bool               `json:"debug"`
	Secret       string             `json:"secret"`
	Host         string             `json:"host"`
	Port         int                `json:"port"`
	BaseUrl      string             `json:"base_url"`
	LogsDatabase ClickhouseDatabase `json:"logs_db"`
	MainDatabase MysqlDatabase      `json:"main_db"`
	Logs         Log                `json:"logs"`
	StoragePath  string             `json:"storage_path"`
}

type MysqlDatabase struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Base     string `json:"base"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ClickhouseDatabase struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Base     string `json:"base"`
	Username string `json:"username"`
	Password string `json:"password"`
}
