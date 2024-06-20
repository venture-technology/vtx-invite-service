package models

type Driver struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	CNH   string `json:"cnh"`
}

type School struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	CNPJ  string `json:"cnpj"`
	Email string `json:"email"`
}

type Invite struct {
	ID     int    `json:"id"`
	School School `json:"school"`
	Driver Driver `json:"driver"`
	Status string `json:"status"`
}
