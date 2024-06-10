package main

import (
	"bufio"
	"os"
	"strings"
)

var (
	DATABASE_URL     string
	DATABASE_API_KEY string
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
}
