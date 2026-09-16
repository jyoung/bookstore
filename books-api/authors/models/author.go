package models

type Author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var AUTHORS = []Author{
	{ID: "1", Name: "Bob"},
	{ID: "2", Name: "Bill"},
}
