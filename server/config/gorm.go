package config

type Mysql struct {
	Dsn          string `mapstructure:"dsn" json:"dsn" yaml:"dsn"`
	//Deprecated: use Dsn
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	//Deprecated: use Dsn
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	//Deprecated: use Dsn
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	//Deprecated: use Dsn
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	//Deprecated: use Dsn
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
}
