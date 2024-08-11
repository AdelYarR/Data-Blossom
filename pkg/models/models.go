package models

type Languages struct {
	ID     int `json:"id"`
	Name   string `json:"name"`
	TypeID int `json:"type_id"`
}

type Types struct {
	ID   int
	Type string
}
