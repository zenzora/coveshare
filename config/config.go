package config

import (
	"encoding/base64"
	"errors"
	"github.com/zenzora/coveshare/secrets"
	"gopkg.in/yaml.v2"
	"os"
)

// GenerateDefaultConfigFile generates a default config file including creating a new master key
func GenerateDefaultConfigFile(path string) error {
	// Check if file exists
	if _, err := os.Stat(path); err == nil {
		return errors.New("config file already exists")
	} else if os.IsNotExist(err) {} else {return err}

	key := secrets.GenerateNewKey()

	f, err := os.Create(path)
	if err != nil {return err}
	defer f.Close()

	encoder := yaml.NewEncoder(f)
	err = encoder.Encode(map[string]string{"key": base64.StdEncoding.EncodeToString(key)})
	if err != nil {return err}
	err = encoder.Close()
	if err != nil {return err}

	return nil
}

