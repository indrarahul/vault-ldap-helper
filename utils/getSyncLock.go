package utils

import (
	"log"
	"vault_ldap_helper/models"
)

//GetSyncLock function for fetching sync lock from vault kv secrets
func GetSyncLock() {

	var sycnLockData models.SyncLock

	GetVaultInfo(Config.Vault.GetSyncLockAPI, &sycnLockData)

	SyncLock = sycnLockData.Data.Lock

	if Verbose > 1 {
		log.Printf("Current SyncLock value : %d", SyncLock)
	}
}
