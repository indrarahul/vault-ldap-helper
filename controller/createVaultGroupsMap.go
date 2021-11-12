package controller

import (
	"log"
	"vault_ldap_helper/models"
	"vault_ldap_helper/utils"
)

//CreateVaultGroupsMap function for storing all usernames of vault groups in VaultGroupsMap
func CreateVaultGroupsMap() {

	var vaultGroupsData models.VaultGroups
	vaultGroupsNameToIDMap := make(map[string]string)
	utils.VaultGroupDataMap = make(map[string][]interface{})

	utils.GetVaultInfo(utils.Config.Vault.GetVaultGroupsAPI, &vaultGroupsData)

	for _, each := range vaultGroupsData.Data.Keys {
		vaultGroupsNameToIDMap[vaultGroupsData.Data.KeyInfo[each].Name] = each
	}

	for _, each := range utils.Config.VaultGroupsList {
		if val, ok := vaultGroupsNameToIDMap[each.Name]; ok {
			var vaultGroupByIDData models.VaultGroupByID
			entityListMap := map[string]int{}

			utils.GetVaultInfo(utils.Config.Vault.GetVaultGroupByIDAPI+val, &vaultGroupByIDData)
			for _, entity := range vaultGroupByIDData.Data.MemberEntityIds {
				entityListMap[entity] = 1
			}

			utils.VaultGroupDataMap[each.Name] = append(utils.VaultGroupDataMap[each.Name], vaultGroupByIDData)
			utils.VaultGroupDataMap[each.Name] = append(utils.VaultGroupDataMap[each.Name], entityListMap)
		}
	}

	if utils.Verbose > 1 {
		log.Printf("VaultGroupDataMap : %v", utils.VaultGroupDataMap)
	}
}
