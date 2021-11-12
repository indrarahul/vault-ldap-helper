package utils

import (
	"log"
	"vault_ldap_helper/models"
)

//UpdateSyncLock function for updating sync lock from vault kv secrets
func UpdateSyncLock(val int) {

	var data models.SyncLock
	data.Data.Lock = val
	PostVaultInfo(Config.Vault.UpdateSyncLockAPI, data.Data)

	if Verbose > 1 {
		log.Printf("Update SyncLock value : %d", val)
	}
}
