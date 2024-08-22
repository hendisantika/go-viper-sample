package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func initializeViper() {
	viper.SetConfigName("config")   // name of config file (without extension)
	viper.SetConfigType("json")     // REQUIRED if the config file does not have the extension
	viper.AddConfigPath("./config") // path to look for the config file in
	err := viper.ReadInConfig()     // Find and read the config file

	if err != nil { // Handle errors reading the config file
		log.Fatalf("Error while reading config file %s", err)
	}
}

func main() {
	initializeViper()

	hostname := viper.GetString("hostname")
	port := viper.GetInt("port")
	username := viper.GetString("credentials.username")
	password := viper.GetString("credentials.password")

	fmt.Printf("Hostname: %s\n", hostname)
	fmt.Printf("Port: %d\n", port)
	fmt.Printf("Username: %s\n", username)
	fmt.Printf("Password: %s\n", password)
}
