package models

type aliases []struct {
	ID            string `json:"id"`
	MountAccessor string `json:"mount_accessor"`
	MountPath     string `json:"mount_path"`
	MountType     string `json:"mount_type"`
	Name          string `json:"name"`
}

type aliasData struct {
	Aliases aliases `json:"aliases"`
}

//VaultEntities data struct
type VaultEntities struct {
	Data struct {
		KeyInfo map[string]aliasData `json:"key_info"`
		Keys    []string             `json:"keys"`
	} `json:"data"`
}
