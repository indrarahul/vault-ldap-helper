# Vault LDAP Helper

Vault LDAP helper works by listing all the members from a LDAP group and adding those members into a Vault group as mentioned in the config below. So you specify group name to a groupFilter(LDAP query).

This is important to note that a user has to be logged in once into Vault before he will be added to the Vault group by this utility. 

This process runs every 3hrs. as it is also configurable from the timeInterval variable in config.yaml

## Build

`make`

## Run

`vault_ldap_helper -config=<CONFIG-FILE-PATH>`

## Configs

Create a config.yaml file and pass this file as an argument as stated above. The config has been defined in detail below.

```
vault:
  url: ""                                               : VAULT SERVERS URL
  getEntitiesAPI: "/v1/identity/entity/id?list=true"    : KEEP THEM AS IT IS UNLESS VAULT API CHANGES
  getVaultGroupsAPI: "/v1/identity/group/id?list=true"  : KEEP THEM AS IT IS UNLESS VAULT API CHANGES
  getVaultGroupByIDAPI: "/v1/identity/group/id/"        : KEEP THEM AS IT IS UNLESS VAULT API CHANGES
  updateVaultGroupByIDAPI: "/v1/identity/group/id/"     : KEEP THEM AS IT IS UNLESS VAULT API CHANGES
  getSyncLockAPI: "/v1/kv/vault_ldap_helper"            : *1
  updateSyncLockAPI: "/v1/kv/vault_ldap_helper"         : *2
  token: ""                                             : VAULT ROOT TOKEN *3
  httpTimeout: 3

timeInterval: 3                                         : *4
ldapCacheExpiration: 3                                  : *5

groups:                                                 : *6
  - 
    name: "group1"
    groupFilter: ""
  - 
    name: "group2"
    groupFilter: ""

ldap:
  hosts:                                                : LDAP SERVERS HOST NAMES
    - "ldap1.com"
    - "ldap2.com"
    - "ldap3.com"
  port: 389                                             : LDAP PORT
  base: ""                                              : LDAP BASE VALUE


  *1 - We added sync lock concept to get rid of race condition between multiple running instances of this process. We do it by storing a key-value pair lock in the vault itself. This api is for getting that lock value and *2 is for updating the lock value. (It's a binary lock having value 0 & 1).

  *3 - As creating and updating a group and adding removing members in a group is root operation so we require root token for it. Decide with the team how you distribute root token to this process and keep it safe.

  *4 - This value corresponds to time interval (in hrs) at which this utility repeats it's processes. As 
        here its 3 so it repeats every 3 hrs. 

  *5 - We maintain cache for ldap information. So, this is the expiration time for that cache.

  *6 - The groups section is a mapping. 

        'name' corresponds to the Group name in the Vault.
        'groupFilter' corresponds to the LDAP query for listing all members in a group. 

        So once you have all the members from the LDAP group those members will be added in the group in 
        Vault given those members have logged in once into Vault.  
```