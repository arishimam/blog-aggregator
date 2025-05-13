package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(user_name string) {
	c.CurrentUserName = user_name
	// call write func to save config struct to json file
	write(*c)

}

func write(cfg Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		fmt.Println("Error getting config file path")
		log.Fatal(err)
		return err
	}
	jsonData, err := json.MarshalIndent(cfg, "", "  ")
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing file")
		log.Fatal(err)
		return err
	}

	return nil
}

func Read() Config {
	file_path, err := getConfigFilePath()
	if err != nil {
		fmt.Println("Error getting config file path")
		log.Fatal(err)
		return Config{}
	}
	file, err := os.Open(file_path)
	if err != nil {
		fmt.Println("Error opening file")
		log.Fatal(err)
		return Config{}
	}
	defer file.Close()

	configInfo := Config{}
	body, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file")
		log.Fatal(err)
		return Config{}
	}
	fmt.Println(string(body))

	err = json.Unmarshal(body, &configInfo)
	if err != nil {
		fmt.Println("Error unmarshaling JSON")
		log.Fatal(err)
		return Config{}
	}

	fmt.Println("DbURL = ", configInfo.DbURL)
	fmt.Println("CurrentUserName", configInfo.CurrentUserName)
	return configInfo
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	filePath := homeDir + "/" + configFileName

	fmt.Println("filePath: ", filePath)
	return filePath, nil
}
