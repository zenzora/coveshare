package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateDefaultConfigFile(t *testing.T) {
	dummyConfigPath := filepath.FromSlash("../bin/testconfig.yaml")

	// Clean up any leftovers
	err := removeIfExists(dummyConfigPath)
	if err != nil {	t.Error(err)}

	// Try regular
	err = GenerateDefaultConfigFile(dummyConfigPath)
	if err != nil {	t.Errorf("Failed to create config file: %s",err)}

	// Attempt to parse
	//@todo

	// Try with config already existing
	err = GenerateDefaultConfigFile(dummyConfigPath)
	if err.Error() != "config file already exists" {
		if err != nil {	t.Error("Method does not return error when attempting overwrite")}
	}

	// Clean up
	err = removeIfExists(dummyConfigPath)
	if err != nil {	t.Error(err)}
}

func removeIfExists(filePath string) error{
	if _, err := os.Stat(filePath); err == nil {
		err := os.Remove(filePath)
		if err != nil {
			return err
		}
	} else if os.IsNotExist(err) {
		return nil

	} else {
		return err
	}
	return nil
}