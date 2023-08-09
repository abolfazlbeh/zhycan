package db

type Config struct {
	SkipDefaultTransaction                   bool `json:"skip_default_transaction"`
	DryRun                                   bool `json:"dry_run"`
	PrepareStmt                              bool `json:"prepare_stmt"`
	DisableAutomaticPing                     bool `json:"disable_automatic_ping"`
	DisableForeignKeyConstraintWhenMigrating bool `json:"disable_foreign_key_constraint_when_migrating"`
	IgnoreRelationshipsWhenMigrating         bool `json:"ignore_relationships_when_migrating"`
	DisableNestedTransaction                 bool `json:"disable_nested_transaction"`
}

type Sqlite struct {
	FileName     string            `json:"db"`
	Options      map[string]string `json:"options"`
	Config       *Config           `json:"config"`
	LoggerConfig *LoggerConfig     `json:"logger"`
}

type Mysql struct {
	DatabaseName string            `json:"db"`
	Username     string            `json:"username"`
	Password     string            `json:"password"`
	Host         string            `json:"host"`
	Port         string            `json:"port"`
	Protocol     string            `json:"protocol"`
	Options      map[string]string `json:"options"`
	Config       *Config           `json:"config"`
	LoggerConfig *LoggerConfig     `json:"logger"`
}

type Postgresql struct {
	DatabaseName string            `json:"db"`
	Username     string            `json:"username"`
	Password     string            `json:"password"`
	Host         string            `json:"host"`
	Port         string            `json:"port"`
	Options      map[string]string `json:"options"`
	Config       *Config           `json:"config"`
	LoggerConfig *LoggerConfig     `json:"logger"`
}

type SqlConfigurable interface {
	Sqlite | Mysql | Postgresql
}

type LoggerConfig struct {
	SlowThreshold             int64  `json:"slow_threshold"`
	IgnoreRecordNotFoundError bool   `json:"ignore_record_not_found_error"`
	ParameterizedQueries      bool   `json:"parameterized_queries"`
	LogLevel                  string `json:"log_level"`
}
