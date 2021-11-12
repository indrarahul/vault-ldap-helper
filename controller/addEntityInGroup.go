package controller

import (
	"log"
	"vault_ldap_helper/models"
	"vault_ldap_helper/utils"
)

//Function for retreiving the LDAP group of a user
func getGroupFromLDAPMap(username string) string {

	for k, v := range LCache.LDAPGroupUsersMap {
		if _, ok := v[username]; ok {

			if utils.Verbose > 2 {
				log.Printf("User %s is in %s LDAP Group", username, k)
			}
			return k
		}
	}

	return ""
}

//helper function for updating the vault group members
func helper(entityName, entityID, ldapGroupName string) {
	var payloadData models.VaultGroupByID

	if val, typeOk := utils.VaultGroupDataMap[ldapGroupName][0].(models.VaultGroupByID); typeOk {
		payloadData = val
	}

	payloadData.Data.MemberEntityIds = append(payloadData.Data.MemberEntityIds, entityID)

	utils.GetSyncLock()

	if utils.SyncLock == 0 { //Checking if lock is acquired
		return
	}

	utils.UpdateSyncLock(0) //Acquire the lock before updating data in vault storage
	utils.UpdateVaultInfo(utils.Config.Vault.UpdateVaultGroupByIDAPI+payloadData.Data.ID, payloadData.Data)
	utils.UpdateSyncLock(1) //Releasing the lock for other instances

	if utils.Verbose > 1 {
		log.Printf("User %s is successfully added into %s vault group", entityName, ldapGroupName)
	}
}

//AddEntityInGroup function for adding an entity to it's corresponding vault group
func AddEntityInGroup() {

	for entityID, entityName := range utils.VaultEntitiesMap {
		ldapGroupOfEntity := getGroupFromLDAPMap(entityName)
		if utils.Verbose > 1 {
			log.Printf("Group of %s: %s", entityName, ldapGroupOfEntity)
		}
		if groupData, ok := utils.VaultGroupDataMap[ldapGroupOfEntity]; ok {
			if val, typeOk := groupData[1].(map[string]int); typeOk {
				if _, entityFound := val[entityID]; entityFound {
					if utils.Verbose > 1 {
						log.Printf("User %s is already in %s: skipping...", entityName, ldapGroupOfEntity)
					}
					continue
				}
				helper(entityName, entityID, ldapGroupOfEntity)
				CreateVaultGroupsMap() // Update the VaultGroupMap after adding users
			}
		}
	}
}
