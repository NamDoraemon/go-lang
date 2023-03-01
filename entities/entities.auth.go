package entities

type Token struct {
	Token string `json:"token"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseLoginFMData struct {
	AccountId              string
	OnesignalExtUserIdHash string
	Token                  string
}

type ResponseLoginFM struct {
	Ok   bool
	Data ResponseLoginFMData
}
