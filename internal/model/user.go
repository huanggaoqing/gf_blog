package model

type RegisterInput struct {
	UserId      string
	UserName    string
	Avatar      string
	PhoneNumber string
	Password    string
	Password2   string
}

type RegisterOutput struct {
	UserId string
}

type LoginInput struct {
}

type LoginOutput struct {
}
