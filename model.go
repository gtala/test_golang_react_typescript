package main

type User struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

type Account struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name"`
	Alias string `json:"alias,omitempty"`
	Owner User   `json:"owner,omitempty"`
}

type Transaction struct {
	ID          string  `json:"id,omitempty"`
	Amount      string  `json:"amount"`
	Source      Account `json:"source,omitempty"`
	Destination Account `json:"destination,omitempty"`
	DateTime    string  `json:"dateTime,omitempty"`
}
