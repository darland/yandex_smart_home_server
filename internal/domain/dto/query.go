package dto

type QueryRequest struct {
	Devices []struct {
		ID         string                 `json:"id"`
		CustomData map[string]interface{} `json:"custom_data,omitempty"`
	} `json:"devices"`
}
