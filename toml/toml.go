package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func initializeViperFromToml2() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("toml")   // Set the type of the configuration file
	viper.AddConfigPath("./toml") // path to look for the config file
	err := viper.ReadInConfig()

	if err != nil { // Handle errors reading the config file
		log.Fatalf("Error while reading config file %s", err)
	}
}

func main() {
	initializeViperFromToml2()
	fmt.Printf("initializeViperFromToml ...")
	// Getting values from the configuration file
	dbServer := viper.GetString("database.server")
	dbPorts := viper.GetIntSlice("database.ports")
	dbConnMax := viper.GetInt("database.connection_max")
	dbEnabled := viper.GetBool("database.enabled")

	// Printing the values
	fmt.Printf("DB Server: %s\n", dbServer)
	fmt.Printf("DB Ports: %v\n", dbPorts)
	fmt.Printf("DB Connection Max: %v\n", dbConnMax)
	fmt.Printf("DB Enabled: %v\n", dbEnabled)
}
