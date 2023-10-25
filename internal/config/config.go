package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/lemjoe/md-blog/internal/models"
)

func LookupAndParseEnvInt(envName string) (int, bool) {
	env, exists := os.LookupEnv(envName)
	if !exists {
		return 0, false
	}
	parsedInt, err := strconv.Atoi(env)
	if err != nil {
		fmt.Printf("warn: %s\n", fmt.Errorf("env '%s' not valid: %w", envName, err))
		return 0, false
	}
	return parsedInt, true
}
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
func InitConfig(confPath string) (models.ConfigDB, error) {
	if confPath != "" {
		if fileExists(confPath) { //if from dot env
			if err := godotenv.Load(confPath); err != nil {
				return models.ConfigDB{}, fmt.Errorf("InitConfig: no '%s' file open", confPath)
			}

		} else {
			return models.ConfigDB{}, fmt.Errorf("InitConfig: '%s' file not exist", confPath)
		}

	}
	defaultConf := models.ConfigDB{
		DbType: "cloverdb",
		Path:   "./db",
	}
	DB_TYPE, exist := os.LookupEnv("DB_TYPE")
	if !exist {
		fmt.Printf("warn: %s\n", fmt.Errorf("env '%s' not found", "DB_TYPE"))
	} else {
		defaultConf.DbType = DB_TYPE
	}

	PORT, exist := os.LookupEnv("PORT")
	if !exist {
		fmt.Printf("warn: %s\n", fmt.Errorf("env '%s' not found", "PORT"))
	} else {
		defaultConf.Port = PORT
	}
	HOST, exist := os.LookupEnv("HOST")
	if !exist {
		fmt.Printf("warn: %s\n", fmt.Errorf("env '%s' not found", "HOST"))
	} else {
		defaultConf.Host = HOST
	}
	NAME, exist := os.LookupEnv("NAME")
	if !exist {
		fmt.Printf("warn: %s\n", fmt.Errorf("env '%s' not found", "NAME"))
	} else {
		defaultConf.Name = NAME
	}
	return defaultConf, nil
}
