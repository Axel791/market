package api

type User struct {
	ID       int64  `json:"id,omitempty"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}
