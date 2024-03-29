package utils

import (
	"os"
	"regexp"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

const projectDirName = "server"

func LoadEnv() {
	if os.Getenv("env") != "HEROKU" {
		re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
		cwd, _ := os.Getwd()
		rootPath := re.Find([]byte(cwd))

		err := godotenv.Load(string(rootPath) + `/.env`)
		if err != nil {
			log.WithFields(log.Fields{
				"cause": err,
				"cwd":   cwd,
			}).Fatal("Problem loading .env file")

			os.Exit(-1)
		}
	}
}
