package models

type MatchRequest struct {
	PetID         string `json:"pet_id"`
	Reason        string `json:"reason"`
	PreferredDay  string `json:"preferred_day"`
	PreferredTime string `json:"preferred_time"`
}
