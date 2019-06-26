package utils

import "os"

func GetDBInfo() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("DB_PASS")
	sslmode := "require"

	return "host=" + host + " port=" + port + " user=" + user + " dbname=" + dbname + " password=" + pass + " sslmode=" + sslmode
}
