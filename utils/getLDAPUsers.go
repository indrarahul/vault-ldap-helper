package utils

import (
	"log"

	"gopkg.in/ldap.v2"
)

//GetLdapUsers function for fetching all usernames from LDAP in particular group
func GetLdapUsers(groupFilter string) *ldap.SearchResult {

	client := &LDAPClient{
		Hosts:       Config.Ldap.Hosts,
		Port:        Config.Ldap.Port,
		Base:        Config.Ldap.Base,
		GroupFilter: groupFilter,
	}

	defer client.Close()

	usersList, err := client.GetUsers()

	if err != nil {
		log.Fatalf("Couldn't find Group :%v", err)
	}

	return usersList
}
