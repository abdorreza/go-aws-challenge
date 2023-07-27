package model

type Device struct {
	Id          string `json:"Id"`
	DeviceModel string `json:"DeviceModel"`
	Name        string `json:"Name"`
	Note        string `json:"Note"`
	Serial      string `json:"Serial"`
}
