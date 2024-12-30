package models

type PostgreSQLStatus struct {
	Active bool `json:"active"`
}

type PostgreSQLMessage struct {
	Message string `json:"message"`
}
