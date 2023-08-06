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
	FileName string            `json:"db"`
	Options  map[string]string `json:"options"`
	Config   *Config           `json:"config"`
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
}

type Postgresql struct {
	DatabaseName string            `json:"db"`
	Username     string            `json:"username"`
	Password     string            `json:"password"`
	Host         string            `json:"host"`
	Port         string            `json:"port"`
	Options      map[string]string `json:"options"`
	Config       *Config           `json:"config"`
}

type SqlConfigurable interface {
	Sqlite | Mysql | Postgresql
}
