package config

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func InitEnvReader() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.AddConfigPath("path")
}

func getEnvValue(key string) string {
	a := viper.Get(key)

	return a.(string)
}

var (
	DbHost = func() string {
		return getEnvValue("DB_HOST")
	}

	DbUser = func() string {
		return getEnvValue("DB_USER")
	}

	DbPassword = func() string {
		return getEnvValue("DB_PASSWORD")
	}

	DbName = func() string {
		return getEnvValue("DB_NAME")
	}

	DbPort = func() string {
		return getEnvValue("DB_PORT")
	}

	DbSslMode = func() string {
		return getEnvValue("DB_SSLMODE")
	}

	DbTimezone = func() string {
		return getEnvValue("DB_TIMEZONE")
	}

	JwtSigningMethod = jwt.SigningMethodHS256

	JwtSignatureKey = func() []byte {
		sk := []byte(getEnvValue("JWT_SIGNATURE_KEY"))
		return []byte(sk)

	}

	ApplicationName = func() string {
		return getEnvValue("APPLICATION_NAME")
	}

	LoginExpirationDuration = 1 * time.Hour

	BcryptCost = func() int {
		costString := getEnvValue("BCRYPT_COST")
		cost, _ := strconv.Atoi(costString)

		return cost
	}

	HttpRequestTimeoutSeconds = func() int {
		costString := getEnvValue("HTTP_REQUEST_TIMEOUT_SECONDS")
		cost, _ := strconv.Atoi(costString)
		return cost
	}
)
