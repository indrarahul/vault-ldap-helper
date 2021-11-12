package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

//UpdateVaultInfo function for updating info on vault server
func UpdateVaultInfo(apiEndpoint string, data interface{}) {

	var headers [][]string
	bearer := fmt.Sprintf("Bearer %s", Config.Vault.Token)
	h := []string{"Authorization", bearer}
	headers = append(headers, h)
	h = []string{"Content-Type", "application/json"}
	headers = append(headers, h)
	h = []string{"Accept", "application/json"}
	headers = append(headers, h)

	apiurl := fmt.Sprintf("%s%s", Config.Vault.URL, apiEndpoint)
	if Verbose > 0 {
		log.Println(apiurl)
	}

	updatedData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Unable to parse Data for Updation, update failed !, error: %v", err)
	}

	req, err := http.NewRequest("PUT", apiurl, bytes.NewBuffer(updatedData))
	if err != nil {
		log.Fatalf("Unable to make request to %s, error: %s", apiurl, err)
	}
	for _, v := range headers {
		if len(v) == 2 {
			req.Header.Add(v[0], v[1])
		}
	}
	if Verbose > 1 {
		dump, err := httputil.DumpRequestOut(req, true)
		if err == nil {
			log.Println("request: ", string(dump))
		}
	}
	timeout := time.Duration(Config.Vault.HTTPTimeout) * time.Second
	client := &http.Client{Timeout: timeout}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Unable to get response from %s, error: %s", apiurl, err)
	}
	if Verbose > 1 {
		dump, err := httputil.DumpResponse(resp, true)
		if err == nil {
			log.Println("response:", string(dump))
		}
	}

	if resp.StatusCode != 204 {
		log.Fatalf("HTTP Respose error, code: %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Unable to read JSON Data from  PUT API, error: %v\n", err)
	}

	if Verbose > 1 {
		log.Println("response Status:", resp.Status)
		log.Println("response Headers:", resp.Header)
		log.Println("response Body:", string(body))
	}
}
