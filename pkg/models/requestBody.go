package models

/* Request Body fields for creation in db */

type RequestBody struct {
	Host_id    int    `json:"host_id"`
	Image_name string `json:"image_name"`
}
