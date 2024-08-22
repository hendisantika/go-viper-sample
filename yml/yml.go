package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func initializeViper() {
	// Set the file name of the configurations file
	viper.SetConfigName("config") // name of config file (without extension)

	// Set the type of the configuration file
	viper.SetConfigType("yml")

	// Set the path to look for the configurations file
	viper.AddConfigPath("./yml") // path to look for the config file in

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil { // Handle errors reading the config file
		log.Fatalf("Error while reading config file %s", err)
	}
}

func main() {
	initializeViper()

	// Getting values from the configuration file
	hostname := viper.GetString("hostname")
	port := viper.GetInt("port")
	username := viper.GetString("credentials.username")
	password := viper.GetString("credentials.password")

	// Printing the values
	fmt.Printf("Hostname: %s\n", hostname)
	fmt.Printf("Port: %d\n", port)
	fmt.Printf("Username: %s\n", username)
	fmt.Printf("Password: %s\n", password)
}
