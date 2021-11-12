package controller

import (
	"log"
	"time"
	"vault_ldap_helper/utils"
)

//LDAPCache - a cache for storing LDAP details and Expiration time for updating the cache
type LDAPCache struct {
	LDAPGroupUsersMap map[string]map[string]int //LDAPGroupUsersMap for storing all usernames in a group from LDAP server
	Expiration        time.Time
}

//UpdateLDAPCache - function for updating the dashboards cache on expiration
func (ldapCache *LDAPCache) UpdateLDAPCache() {

	if !utils.FirstRunSinceRestart && ldapCache.Expiration.After(time.Now()) {
		return
	}

	createLDAPGroupUsersMap()
	if utils.Verbose > 1 {
		log.Printf("Succesfully Updated the LDAPCache\n")
	}
	ldapCache.Expiration = time.Now().Add(utils.Config.LDAPCacheExpiration * time.Hour)
}

//LCache - variable for LDAPCache
var LCache LDAPCache

//createLDAPGroupUsersMap function for storing all usernames in in LDAP Groups
func createLDAPGroupUsersMap() {

	var tmp map[string]int
	LCache.LDAPGroupUsersMap = make(map[string]map[string]int)

	for _, each := range utils.Config.VaultGroupsList {
		ldapSearchRes := utils.GetLdapUsers(each.GroupFilter)
		tmp = make(map[string]int)
		for _, entry := range ldapSearchRes.Entries {
			username := entry.GetAttributeValue("cn")
			if utils.Verbose > 1 {
				log.Printf("User %s found in %s LDAP Group\n", username, each)
			}
			tmp[username] = 1
		}

		if utils.Verbose > 1 {
			log.Printf("User Map for %s Group :\n %v\n", each.Name, tmp)
		}

		LCache.LDAPGroupUsersMap[each.Name] = tmp
	}

	if utils.Verbose > 1 {
		log.Printf("LDAPGroupUsersMap : %v", LCache.LDAPGroupUsersMap)
	}
}
