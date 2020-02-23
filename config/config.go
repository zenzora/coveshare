package config

import (
	"encoding/base64"
	"errors"
	"os"

	"github.com/spf13/viper"
	"github.com/zenzora/coveshare/secrets"
	"gopkg.in/yaml.v2"
)

// GenerateDefaultConfigFile generates a default config file including creating a new master key
func GenerateDefaultConfigFile(path string) error {
	// Check if file exists
	if _, err := os.Stat(path); err == nil {
		return errors.New("config file already exists")
	} else if os.IsNotExist(err) {
	} else {
		return err
	}

	key := secrets.GenerateNewKey()

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := yaml.NewEncoder(f)
	err = encoder.Encode(map[string]string{
		"encryption-type": "aes-sha256",
		"key":             base64.StdEncoding.EncodeToString(key[:]),
	})

	if err != nil {
		return err
	}
	err = encoder.Close()
	if err != nil {
		return err
	}

	return nil
}

//Validate checks if the config is setup correctly
func Validate() error {
	// Validate the encryption type
	encryptionType := viper.GetString("encryption-type")
	if encryptionType == "aes-sha256" {
		key, _ := base64.StdEncoding.DecodeString(viper.GetString("key"))
		if len(key) != 32 {
			return errors.New("key must be a 32 byte string ")
		}
	} else {
		return errors.New("encryption type must be \"aes-sha256\"")
	}
	return nil
}
