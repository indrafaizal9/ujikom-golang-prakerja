package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"ujikom/database"

	"github.com/spf13/viper"
)

// Path: cmd/cli/main.go
func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: go run cmd/cli/main.go [command]\n\n")
		fmt.Println("Available commands:")
		fmt.Println("	- key:generate")
		fmt.Println("	- db:seed")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "key:generate":
		generateKey()
	case "db:seed":
		//if seed name is not provided, run all seeder
		if len(os.Args) < 3 {
			Seed("all")
		} else {
			Seed(os.Args[2])
		}
		fmt.Println("Seeding completed!")
	default:
		fmt.Println("Invalid command")
		os.Exit(1)
	}
}

func generateKey() {
	randomstring, err := generateRandomString(32)
	if err != nil {
		fmt.Println("Error generating random string")
		os.Exit(1)
	}
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName("app")
	viper.AutomaticEnv()
	viper.SetDefault("APP_SECRET", randomstring)
	viper.WriteConfig()
	// viper.Set("APP_SECRET", randomstring)
	// viper.WriteConfigAs("app.env")
}

func generateRandomString(length int) (string, error) {
	bytesNeeded := (length * 3) / 4

	randomBytes := make([]byte, bytesNeeded)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	randomString := base64.URLEncoding.EncodeToString(randomBytes)
	randomString = randomString[:length]

	return randomString, nil
}

func Seed(name string) {
	switch name {
	case "all":
		SeederTagsAndLabels()
		SeederUsers()
	case "tags_and_labels":
		SeederTagsAndLabels()
	case "users":
		SeederUsers()
	default:
		fmt.Println("Invalid seeder name")
	}
}

func SeederTagsAndLabels(){
	Seeder := database.Seeder{}
	Seeder.SeederTagsAndLabels()
}

func SeederUsers(){
	Seeder := database.Seeder{}
	Seeder.SeederUsers()
}