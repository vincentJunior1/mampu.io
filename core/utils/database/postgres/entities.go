package postgres

type ConfigPostgres struct {
	Username string
	Password string
	Host     string
	Port     string
	DbName   string
	Debug    string
	Mode     string
}
