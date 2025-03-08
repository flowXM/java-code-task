package configs

import (
	"fmt"
	"os"
)

var (
	DBUser     = os.Getenv("POSTGRES_USER")
	DBPassword = os.Getenv("POSTGRES_PASSWORD")
	DBName     = os.Getenv("POSTGRES_DB")
	DBHost     = os.Getenv("POSTGRES_SERVER")
	DBPort     = os.Getenv("POSTGRES_PORT")
)

func GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", DBHost, DBPort, DBUser, DBName, DBPassword)
}
