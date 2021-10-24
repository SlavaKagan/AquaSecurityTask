package models

// container fields

type Container struct {
	ID         int    `json:"id"`
	Host_ID    int    `json:"host_id"`
	Name       string `json:"name"`
	Image_Name string `json:"image_name"`
	Host_name  string `json:"host_name"`
}
