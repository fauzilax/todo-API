package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

// Deklarasi Variable Global Untuk Memanggil file Secret Key di Env
var (
	JWTKey = ""
)

type DBConfig struct {
	Mysql_User     string
	Mysql_Password string
	Mysql_Host     string
	Mysql_Port     int
	Mysql_DBName   string
	jwtKey         string
}

// membuat fungsi global untuk pemanggilan config
func InitConfig() *DBConfig {
	return ReadEnv()
}

func ReadEnv() *DBConfig {
	app := DBConfig{}
	isRead := true

	if val, found := os.LookupEnv("MYSQL_USER"); found {
		app.Mysql_User = val
		isRead = false
	}
	if val, found := os.LookupEnv("MYSQL_PASSWORD"); found {
		app.Mysql_Password = val
		isRead = false
	}
	if val, found := os.LookupEnv("MYSQL_HOST"); found {
		app.Mysql_Host = val
		isRead = false
	}
	if val, found := os.LookupEnv("MYSQL_PORT"); found {
		cnv, _ := strconv.Atoi(val)
		app.Mysql_Port = cnv
		isRead = false
	}
	if val, found := os.LookupEnv("MYSQL_DBNAME"); found {
		app.Mysql_DBName = val
		isRead = false
	}
	if val, found := os.LookupEnv("JWT_KEY"); found {
		app.jwtKey = val
		isRead = false
	}

	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
		viper.SetConfigType("env")
		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config : ", err.Error())
			return nil
		}
		err = viper.Unmarshal(&app)
		if err != nil {
			log.Println("error parse config : ", err.Error())
			return nil
		}
	}
	JWTKey = app.jwtKey

	return &app
}
