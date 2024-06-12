package constants

import (
	"bufio"
	"os"
	"strings"
)

var (
	DATABASE_URL     string
	DATABASE_API_KEY string
	PLAID_SECRET     string
	PLAID_CLIENT_ID  string
	CLIENT_NAME      string
	JWT_SECRET       string
)

func loadEnv() {
	file, err := os.Open(".env.local")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			os.Setenv(parts[0], parts[1])
		}
	}
}

func init() {
	loadEnv()

	DATABASE_URL = os.Getenv("DATABASE_API_URL")
	if DATABASE_URL == "" {
		DATABASE_URL = os.Getenv("DATABASE_API_URL_DEV")
	}

	DATABASE_API_KEY = os.Getenv("DATABASE_API_KEY")
	if DATABASE_API_KEY == "" {
		DATABASE_API_KEY = os.Getenv("DATABASE_API_KEY_DEV")
	}

	PLAID_CLIENT_ID = os.Getenv("PLAID_CLIENT_ID")
	if PLAID_CLIENT_ID == "" {
		PLAID_CLIENT_ID = os.Getenv("PLAID_CLIENT_ID_DEV")
	}

	PLAID_SECRET = os.Getenv("PLAID_SECRET")
	if PLAID_SECRET == "" {
		PLAID_SECRET = os.Getenv("PLAID_SECRET_DEV")
	}
	CLIENT_NAME = "NetWorthNavigator"
	JWT_SECRET = os.Getenv("JWT_SECRET")

}
