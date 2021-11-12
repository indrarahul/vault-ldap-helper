package models

import (
	"time"
)

//Config data struct
type Config struct {
	Vault struct {
		URL                     string `yaml:"url"`
		GetEntitiesAPI          string `yaml:"getEntitiesAPI"`
		GetVaultGroupsAPI       string `yaml:"getVaultGroupsAPI"`
		GetVaultGroupByIDAPI    string `yaml:"getVaultGroupByIDAPI"`
		UpdateVaultGroupByIDAPI string `yaml:"updateVaultGroupByIDAPI"`
		GetSyncLockAPI          string `yaml:"getSyncLockAPI"`
		UpdateSyncLockAPI       string `yaml:"updateSyncLockAPI"`
		Token                   string `yaml:"token"`
		HTTPTimeout             int    `yaml:"httpTimeout"`
	} `yaml:"vault"`

	TimeInterval        time.Duration `yaml:"timeInterval"`        //Time interval at which the script will repeats itself (in hrs.)
	LDAPCacheExpiration time.Duration `yaml:"ldapCacheExpiration"` //(in hrs.)

	VaultGroupsList []struct {
		Name        string `yaml:"name"`
		GroupFilter string `yaml:"groupFilter"`
	} `yaml:"groups"`

	Ldap struct {
		Hosts []string `yaml:"hosts"`
		Port  int      `yaml:"port"`
		Base  string   `yaml:"base"`
	} `yaml:"ldap"`
}
