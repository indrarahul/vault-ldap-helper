package controller

import (
	"log"
	"vault_ldap_helper/models"
	"vault_ldap_helper/utils"
)

//CreateVaultEntitiesMap function for storing all usernames in VaultEntitiesMap
func CreateVaultEntitiesMap() {

	var vaultEntitiesData models.VaultEntities

	utils.GetVaultInfo(utils.Config.Vault.GetEntitiesAPI, &vaultEntitiesData)

	utils.VaultEntitiesMap = make(map[string]string)

	for _, each := range vaultEntitiesData.Data.Keys {
		utils.VaultEntitiesMap[each] = vaultEntitiesData.Data.KeyInfo[each].Aliases[0].Name
	}

	if utils.Verbose > 1 {
		log.Printf("VaultEntitiesMap : %v", utils.VaultEntitiesMap)
	}

}
