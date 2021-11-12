package utils

import (
	"io/ioutil"
	"log"
	"os"
	"vault_ldap_helper/models"

	"gopkg.in/yaml.v2"
)

//Verbose variable
var Verbose int

//Config variable
var Config models.Config

//VaultEntitiesMap for storing all usernames in entities folder in vault server
var VaultEntitiesMap map[string]string

//VaultGroupDataMap s
var VaultGroupDataMap map[string][]interface{}

//SyncLock for storing lock which eliminates the race conditions btw multiple instances of vault_ldap_helper on vault cluster
var SyncLock int

//FirstRunSinceRestart - boolean variable for storing the information if the it's the first start of the service after restart or not
var FirstRunSinceRestart bool

//ParseConfig - Function for parsing the config File
func ParseConfig(configFilePath string) {

	if stats, err := os.Stat(configFilePath); err == nil {
		if Verbose > 2 {
			log.Printf("FileInfo: %s\n", stats)
		}

		yamlFile, e := ioutil.ReadFile(configFilePath)
		if e != nil {
			log.Fatalf("Config File can't be read, error: %s", e)
		}

		err := yaml.Unmarshal(yamlFile, &Config)
		if err != nil {
			log.Fatalf("Config JSON File can't be loaded, error: %s", err)
		} else if Verbose > 0 {
			log.Printf("Loading config from %s\n", configFilePath)
		}
	} else {
		log.Fatalf("%s: Config File doesn't exist, error: %v", configFilePath, err)
	}

	if Verbose > 0 {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	} else {
		log.SetFlags(log.LstdFlags)
	}

	if Verbose > 2 {
		log.Printf("Configuration:\n%+v\n", Config)
	}

}
