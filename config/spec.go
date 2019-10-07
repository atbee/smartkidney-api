package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	// Specification holds environment variable name.
	Specification struct {
		DBHost               string
		DBName               string
		DBUsersCol           string
		DBBloodPressureCol   string
		DBGlomerularInfilCol string
		DBBloodSugarCol      string
		DBBMICol             string
		DBWaterCol           string
		APIPort              string
	}
)

// Spec retrieves the value of the environment variable named by the key.
func Spec() *Specification {
	godotenv.Load()

	s := Specification{
		DBHost:               os.Getenv("DB_HOST"),
		DBName:               os.Getenv("DB_NAME"),
		DBUsersCol:           os.Getenv("DB_USERS_COL"),
		DBBloodPressureCol:   os.Getenv("DB_BLOODPRESSURE_COL"),
		DBGlomerularInfilCol: os.Getenv("DB_GLOMERULARINFIL_COL"),
		DBBloodSugarCol:      os.Getenv("DB_BLOODSUGAR_COL"),
		DBBMICol:             os.Getenv("DB_BMI_COL"),
		DBWaterCol:           os.Getenv("DB_WATER_COL"),
		APIPort:              os.Getenv("API_PORT"),
	}
	return &s
}
