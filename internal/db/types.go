package db

type SqliteConfig struct {
	FileName string            `json:"db"`
	Options  map[string]string `json:"options"`
}

type MysqlConfig struct {
	DatabaseName string            `json:"db"`
	Username     string            `json:"username"`
	Password     string            `json:"password"`
	Host         string            `json:"host"`
	Port         string            `json:"port"`
	Protocol     string            `json:"protocol"`
	Options      map[string]string `json:"options"`
}

type PostgresqlConfig struct {
	DatabaseName string            `json:"db"`
	Username     string            `json:"username"`
	Password     string            `json:"password"`
	Host         string            `json:"host"`
	Port         string            `json:"port"`
	Options      map[string]string `json:"options"`
}

type SqlConfigurable interface {
	SqliteConfig | MysqlConfig | PostgresqlConfig
}
