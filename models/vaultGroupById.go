package models

//VaultGroupByID data struct
type VaultGroupByID struct {
	Data struct {
		ID              string   `json:"id"`
		Name            string   `json:"name"`
		Type            string   `json:"type"`
		Policies        []string `json:"policies"`
		MemberEntityIds []string `json:"member_entity_ids"`
	} `json:"data"`
}
