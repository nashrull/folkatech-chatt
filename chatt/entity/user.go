package request

type RegisterUser struct {
	Nohp string `json:"nohp,omitempty"`
}

type RegisterUserResponse struct {
	Id   interface{} `json:"id,omitempty"`
	Nohp string      `json:"nohp,omitempty"`
}
