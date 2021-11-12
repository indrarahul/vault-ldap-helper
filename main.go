package main

import (
	"flag"
	"time"
	"vault_ldap_helper/controller"
	"vault_ldap_helper/utils"
)

func run() {

	for true {

		utils.FirstRunSinceRestart = false
		ptr := &controller.LCache
		ptr.UpdateLDAPCache()

		controller.CreateVaultEntitiesMap()
		controller.CreateVaultGroupsMap()
		controller.AddEntityInGroup()
		time.Sleep(utils.Config.TimeInterval * time.Hour)
	}
}

func main() {

	var configPath string

	flag.StringVar(&configPath, "config", "", "Config File Path")
	flag.IntVar(&utils.Verbose, "verbose", 0, "Verbosity value")
	flag.Parse()

	utils.ParseConfig(configPath)

	run()
}
