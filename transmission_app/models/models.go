package models

type TransmissionStatus struct {
    Active bool `json:"active"`
}

type TransmissionMessage struct {
    Message string `json:"message"`
}