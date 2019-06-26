package utils

import "os"

func GetDBInfo() string {
	host := os.GetEnv("DB_HOST")
	port := os.GetEnv("DB_PORT")
	user := os.GetEnv("DB_USER")
	dbname := os.GetEnv("DB_NAME")
	pass := os.GetEnv("DB_PASS")
	sslmode := "require"

	return "host=" + host + "port=" + port + "user=" + user + "dbname=" + dbname + "password=" + pass + "sslmode=" + sslmode
}
