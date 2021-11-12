package models

type groupData struct {
	Name              string `json:"name"`
	NumMemberEntities int    `json:"num_member_entities"`
	NumParentGroups   int    `json:"num_parent_groups"`
}

//VaultGroups data struct
type VaultGroups struct {
	Data struct {
		KeyInfo map[string]groupData `json:"key_info"`
		Keys    []string             `json:"keys"`
	} `json:"data"`
}
