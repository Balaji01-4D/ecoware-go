package initializer

import	(
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading the .env files")
	}
}