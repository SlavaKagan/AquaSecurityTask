package models

// host represents data
type Host struct {
	ID         int    `json:"id"`
	Uuid       string `json:"uuid"`
	Name       string `json:"name"`
	Ip_address string `json:"ip_address"`
}
