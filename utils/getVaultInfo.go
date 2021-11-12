package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
	"vault_ldap_helper/models"
)

//GetVaultInfo function for fetching info from vault server
func GetVaultInfo(apiEndpoint string, refData interface{}) {

	var data interface{}

	if apiEndpoint == Config.Vault.GetEntitiesAPI {
		data = refData.(*models.VaultEntities)
	} else if apiEndpoint == Config.Vault.GetVaultGroupsAPI {
		data = refData.(*models.VaultGroups)
	} else if strings.Contains(apiEndpoint, Config.Vault.GetVaultGroupByIDAPI) {
		data = refData.(*models.VaultGroupByID)
	} else if apiEndpoint == Config.Vault.GetSyncLockAPI {
		data = refData.(*models.SyncLock)
	}

	var headers [][]string
	bearer := fmt.Sprintf("Bearer %s", Config.Vault.Token)
	h := []string{"Authorization", bearer}
	headers = append(headers, h)
	h = []string{"Accept", "application/json"}
	headers = append(headers, h)

	apiURL := fmt.Sprintf("%s%s", Config.Vault.URL, apiEndpoint)

	if Verbose > 0 {
		log.Println(apiURL)
	}

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		log.Fatalf("Unable to make request to %s, error: %s", apiURL, err)
	}
	for _, v := range headers {
		if len(v) == 2 {
			req.Header.Add(v[0], v[1])
		}
	}

	timeout := time.Duration(Config.Vault.HTTPTimeout) * time.Second
	client := &http.Client{Timeout: timeout}

	if Verbose > 1 {
		log.Println("URL", apiURL)
		dump, err := httputil.DumpRequestOut(req, true)
		if err == nil {
			log.Println("Request: ", string(dump))
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Response Error, error: %v\n", err)
	}
	defer resp.Body.Close()

	byteValue, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Unable to read JSON Data GET API, error: %v\n", err)
	}

	if Verbose > 0 {
		log.Println(string(byteValue))
	}
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		log.Fatalf("Unable to parse JSON Data from GET API, error: %v\n", err)
	}

	if Verbose > 1 {
		dump, err := httputil.DumpResponse(resp, true)
		if err == nil {
			log.Println("Response: ", string(dump))
		}
	}
}
