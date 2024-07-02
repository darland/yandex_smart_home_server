package dto

type Payload struct {
	UserID  string `json:"user_id"`
	Devices []struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Room        string `json:"room"`
		Type        string `json:"type"`
		CustomData  struct {
			Foo  int           `json:"foo"`
			Bar  string        `json:"bar"`
			Baz  bool          `json:"baz"`
			Qux  []interface{} `json:"qux"`
			Quux struct {
				Quuz struct {
					Corge []interface{} `json:"corge"`
				} `json:"quuz"`
			} `json:"quux"`
		} `json:"custom_data"`
		Capabilities []struct {
			Type        string `json:"type"`
			Retrievable bool   `json:"retrievable,omitempty"`
			Parameters  struct {
				Instance string `json:"instance,omitempty"`
				Unit     string `json:"unit,omitempty"`
				Range    struct {
					Min       int `json:"min"`
					Max       int `json:"max"`
					Precision int `json:"precision"`
				} `json:"range,omitempty"`
				ColorModel   string `json:"color_model,omitempty"`
				TemperatureK struct {
					Min       int `json:"min"`
					Max       int `json:"max"`
					Precision int `json:"precision"`
				} `json:"temperature_k,omitempty"`
			} `json:"parameters,omitempty"`
		} `json:"capabilities"`
		DeviceInfo struct {
			Manufacturer string `json:"manufacturer"`
			Model        string `json:"model"`
			HwVersion    string `json:"hw_version"`
			SwVersion    string `json:"sw_version"`
		} `json:"device_info"`
	} `json:"devices"`
}

type DevicesResponse struct {
	RequestID string   `json:"request_id"`
	Payload   *Payload `json:"payload"`
}
