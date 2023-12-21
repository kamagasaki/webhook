package webhook

import "os"

const WAKeyword = "wh4t5auth0"

var (
	WebhookSecret = os.Getenv("SECRET")
	WAAPIToken    = os.Getenv("TOKEN")
)

var (
	DBHostname = os.Getenv("DB_HOST")
	DBUsername = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBPortNum  = os.Getenv("DB_PORT")
	DBName     = os.Getenv("DB_NAME")
)
