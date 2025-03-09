package database

type DB struct {
	Host string
	User string
	Pass string
	Port string
	Name string
}

type DBPostgreSQL struct {
	DB
	SslMode  string
	TimeZone string
}
